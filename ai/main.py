import os
from contextlib import asynccontextmanager

import pandas as pd
from catboost import CatBoostRegressor, Pool
from fastapi import FastAPI, HTTPException
from model import RawBond, PredictRequest
from funcs import transformData

model: CatBoostRegressor = None

@asynccontextmanager
async def lifespan(app: FastAPI):
    global model
    model_path = "learn/model"

    if not os.path.exists(model_path):
        raise FileNotFoundError(f"Файл модели '{model_path}' не найден в корневой директории!")

    model = CatBoostRegressor()
    model.load_model(model_path)

    yield


app = FastAPI(lifespan=lifespan)


@app.post("/predict")
def predict(request: PredictRequest):
    if model is None:
        raise HTTPException(status_code=500, detail="Модель не была загружена на сервере.")

    if not request.bonds:
        return {"predictions": []}

    if request.kr is None or request.kr == 0:
        raise HTTPException(status_code=400, detail="Ключевая ставка не была передана")

    bonds_dicts = [bond.model_dump() for bond in request.bonds]
    df = pd.DataFrame(bonds_dicts)

    df.head()
    df_meta = df[['id', 'price']].copy()

    df = transformData(df, request.kr)
    df_meta = df_meta.loc[df.index]

    cat_features = [
        "currency_id",
        "rating"
    ]

    cat_features = [c for c in cat_features if c in df.columns]

    test_pool = Pool(df, df['price'], cat_features=cat_features)

    predictions = model.predict(test_pool)

    suspicious_ids = []

    THRESHOLD_PERCENT = 0.15

    for idx, (original_idx, row) in enumerate(df_meta.iterrows()):
        real_price = float(row['price'])
        predicted_price = predictions[idx]
        bond_id = row['id']

        if real_price == 0:
            continue

        relative_error = abs(real_price - predicted_price) / real_price

        if relative_error > THRESHOLD_PERCENT:
            suspicious_ids.append(bond_id)

    return {
        "suspicious_bond_ids": suspicious_ids
    }


