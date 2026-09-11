import type { KeyRate, KeyRateRaw, Ruonia } from './model';

const headers = new Headers({
  'Content-Type': 'application/json',
});

export async function GetKeyRate() {
  const resp = await fetch('/backend/api/v1/analitics/key-rate/ru',
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка получения ключевой ставки');
  }

  const stringKeyRate: string = await resp.json();

  return Number.parseFloat(stringKeyRate);
}

async function getRates(url: string, errorMessage: string): Promise<KeyRate[]> {
  const resp = await fetch(url,
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error(errorMessage);
  }

  const keyRates: KeyRateRaw[] = await resp.json();
  const result: KeyRate[] = [];

  for (const elem of keyRates) {
    result.push(
      {
        rate: Number.parseFloat(elem.rate),
        date: new Date(elem.date)
      }
    );
  }

  return result;
}

export function GetKeyRates(): Promise<KeyRate[]> {
  return getRates('/backend/api/v1/analitics/key-rate/ru/full', 'Ошибка получения ключевой ставки');
}

export function GetRuonia(): Promise<Ruonia[]> {
  return getRates('/backend/api/v1/analitics/ruonia/ru/full', 'Ошибка получения ставки RUONIA');
}