export type Bond = {
    id: string
    companyID: string
    isin: string
    name: string
    type: string
    subType: string
    price: number
    ytm: number
    duration: number
    lotSize: number
    faceValue: number
    couponPercent: number
    couponPeriod: number
    nextCoupon: Date
    callOption: Date | null
    putOption: Date | null
    matDate: Date
    valToday: number
    acruedint: number
    issueSize: number
    currencyID: string
    boardID: string
}

export type Rating = {
    id: string
    companyID: string
    ratingValue: string
    agencyName: string
    releaseUrl: string
    objectName: string
    releaseDate: Date
    isRevoked: boolean
}

export type Company = {
    id: string
    name: string
}

export type CompanyWithRating = Company & {
    ratings: Rating[]
}

export type BondWithRatings = Bond & {
    ratings: Rating[]
}

export type FullBonds = {
    bonds: Bond[]
    companies: CompanyWithRating[]
}

const ratingOrder: Record<string, number> = {
  'AAA': 1,

  'AA+': 2, 'AA' : 3, 'AA-': 4,
  'A+' : 5, 'A'  : 6, 'A-' : 7,

  'BBB+': 8, 'BBB' : 9, 'BBB-': 10,
  'BB+' : 11, 'BB'  : 12, 'BB-' : 13,
  'B+'  : 14, 'B'   : 15, 'B-'  : 16,

  'CCC+': 17, 'CCC' : 18, 'CCC-': 19,
  'CC'  : 20,
  'C'   : 21,
  'D'   : 22,
};

export function compareRatings(a: string, b: string): number {
  const rankA = ratingOrder[a] ?? Number.MAX_SAFE_INTEGER;
  const rankB = ratingOrder[b] ?? Number.MAX_SAFE_INTEGER;

  return rankA - rankB;
}

export function inRange(from: string, to: string, target: string): boolean {
  if (target === '') {
    return false;
  }

  const rankFrom = ratingOrder[from] ?? Number.MAX_SAFE_INTEGER;
  const rankTo = ratingOrder[to] ?? Number.MAX_SAFE_INTEGER;
  const rankTarget = ratingOrder[target] ?? Number.MAX_SAFE_INTEGER;
  
  return rankTarget <= rankFrom && rankTarget >= rankTo;
}

export function getBondWithRating(bondID: string, bonds: Bond[], companies: CompanyWithRating[]): BondWithRatings | null {
  const bond = bonds.find(bond => bond.id === bondID);

  if (bond === undefined) {
    return null;
  }

  return {
    ...bond,
    ratings: companies.find(company => company.id === bond.companyID)?.ratings || []
  };
}

export function getLastNotRevokedRatings(companies: CompanyWithRating[]) {
  const map = new Map<string, string>();

  for (const company of companies) {
    for (const rating of company.ratings) {
      if (!rating.isRevoked) {
        map.set(company.id, rating.ratingValue);
        break;
      }
    }
  }

  return map;
}