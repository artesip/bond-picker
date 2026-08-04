import type { BondWithRatings } from '../bonds/model';


const headers = new Headers({
  'Content-Type': 'application/json',
});

export async function GetSuspicuous(bonds: BondWithRatings[], kr: number): Promise<string[]> {
  const bondsBody = bonds.map(bond => ({ ...bond, rating: bond.ratings?.[0]?.ratingValue ?? null }));

  const resp = await fetch('/ai/predict',
    {
      method : 'POST',
      headers: headers,
      body   : JSON.stringify(
        {
          'bonds': bondsBody,
          'kr'   : kr
        }
      )
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка получения рекомендаций');
  }

  const data = await resp.json();

  return data.suspicious_bond_ids;
}