import pandas as pd


def transformData(df: pd.DataFrame, kr: float) -> pd.DataFrame:
    df = df.copy()

    rename_map = {
        "companyID": "company_id", "subType": "sub_type", "lotSize": "lot_size",
        "faceValue": "face_value", "couponPercent": "coupon_percent", "couponPeriod": "coupon_period",
        "nextCoupon": "next_coupon", "callOption": "call_option", "putOption": "put_option",
        "matDate": "mat_date", "valToday": "val_today", "issueSize": "issue_size",
        "currencyID": "currency_id", "boardID": "board_id"
    }
    df = df.rename(columns=rename_map)

    df["kr"] = kr

    columns_to_drop = [
        'id', 'rating_date', 'company_id', 'isin', 'type', 'sub_type',
        'board_id', 'created_at', 'updated_at', 'lot_size', 'name',
    ]
    df = df.drop(columns=columns_to_drop, errors='ignore')

    df = df.dropna(subset=['rating'])

    today = pd.Timestamp.now(tz="UTC")

    df["mat_date"] = pd.to_datetime(df["mat_date"])
    df = df[df["mat_date"] > today]

    df["ttm"] = (df["mat_date"] - today).dt.days / 365
    df = df.drop(columns=['mat_date'])

    df["has_put"] = df["put_option"].notna().astype(int)
    df["has_call"] = df["call_option"].notna().astype(int)
    df = df.drop(columns=['put_option', 'call_option'], errors='ignore')

    df["next_coupon"] = pd.to_datetime(df["next_coupon"])
    df["days_to_coupon"] = (df["next_coupon"] - today).dt.days
    df["coupon_cycle_progress"] = df["days_to_coupon"] / 365
    df["coupon_soon"] = (df["days_to_coupon"] <= 30).astype(int)

    df = df.drop(columns=['next_coupon'])

    columns_order = [
        'currency_id', 'ytm', 'duration', 'val_today', 'face_value',
        'coupon_period', 'coupon_percent', 'issue_size', 'acruedint',
        'rating', 'kr', 'has_put', 'has_call', 'ttm',
        'days_to_coupon', 'coupon_cycle_progress', 'coupon_soon', 'price'
    ]

    return df[columns_order]