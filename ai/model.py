from typing import Optional, List

from pydantic import BaseModel

class RawBond(BaseModel):
    id: str
    companyID: str
    isin: str
    name: str
    type: str
    subType: str
    price: float
    ytm: float
    duration: float
    lotSize: int
    faceValue: float
    couponPercent: float
    couponPeriod: int
    nextCoupon: str
    callOption: Optional[str] = None
    putOption: Optional[str] = None
    matDate: str
    valToday: float
    acruedint: float
    issueSize: float
    currencyID: str
    boardID: str
    rating: Optional[str] = None

class PredictRequest(BaseModel):
    bonds: List[RawBond]
    kr: float