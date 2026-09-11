
export type KeyRateRaw = {
  rate: string
  date: Date
}

export type KeyRate = {
  rate: number
  date: Date
}

export type RuoniaRaw = KeyRateRaw
export type Ruonia = KeyRate