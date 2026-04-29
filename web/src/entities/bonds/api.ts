import type { Bond, Rating } from './model';

const headers = new Headers({
  'Content-Type': 'application/json',
});

function truncateDecimals(num: number, digits: number): number {
  const factor = 10 ** digits;
  return Math.trunc(num * factor) / factor;
}

export async function GetBonds(): Promise<Bond[]> {
  const resp = await fetch('/backend/api/v1/bond?type=fix', 
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка получения облигаций');
  }

  const data: Bond[] = await resp.json();
  const result = data.map((element) => ({
    ...element,
    price        : truncateDecimals(element.price, 2),
    ytm          : truncateDecimals(element.ytm, 2),
    faceValue    : truncateDecimals(element.faceValue, 2),
    couponPercent: truncateDecimals(element.couponPercent, 2),
    acruedint    : truncateDecimals(element.acruedint, 2),
    duration     : truncateDecimals(element.duration / 365, 2),

    nextCoupon: new Date(element.nextCoupon),
    matDate   : new Date(element.matDate),
    callOption: element.callOption ? new Date(element.callOption) : element.callOption,
    putOption : element.putOption ? new Date(element.putOption) : element.putOption
  }))
    .filter((element) => element.ytm <= 150.0 && element.ytm >= 0.0 && element.duration > 0.0);

  return result;
}

export async function GetRatings(): Promise<Rating[]> {
  const resp = await fetch('/backend/api/v1/bond/rating', 
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка получения рейтингов');
  }

  const ratings: Rating[] = await resp.json();
  const result = ratings.map((element) => ({
    ...element,
    releaseDate: new Date(element.releaseDate),
  })).sort((a, b) => b.releaseDate.getTime() - a.releaseDate.getTime());

  return result;
}

export async function PickBond(id: string, count: number) {
  const resp = await fetch(`/backend/api/v1/bond/pick/${id}?count=${count}`,
    {
      method : 'POST',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка добавления в избранное');
  }
}

export async function GetPicked() {
  const resp = await fetch('/backend/api/v1/bond/pick',
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка избранных облигаций');
  }

  const data: Bond[] = await resp.json();
  const result = data.map((element) => ({
    ...element,
    price        : truncateDecimals(element.price, 2),
    ytm          : truncateDecimals(element.ytm, 2),
    faceValue    : truncateDecimals(element.faceValue, 2),
    couponPercent: truncateDecimals(element.couponPercent, 2),
    acruedint    : truncateDecimals(element.acruedint, 2),
    duration     : truncateDecimals(element.duration / 365, 2),

    nextCoupon: new Date(element.nextCoupon),
    matDate   : new Date(element.matDate),
    callOption: element.callOption ? new Date(element.callOption) : element.callOption,
    putOption : element.putOption ? new Date(element.putOption) : element.putOption
  }))
    .filter((element) => element.ytm <= 150.0 && element.ytm >= 0.0 && element.duration > 0.0);

  return result;
}

export async function DeletePicked(id: string) {
  const resp = await fetch(`/backend/api/v1/bond/pick/${id}`,
    {
      method : 'DELETE',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка удаления избранной облигации');
  }
}

