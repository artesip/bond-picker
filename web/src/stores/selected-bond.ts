import { create } from 'zustand';

type SelectedBondState = {
  selectedId: string | null;
  setSelectedId: (id: string | null) => void;
};

export const useSelectedBondStore = create<SelectedBondState>((set) => ({
  selectedId   : null,
  setSelectedId: (selectedId) => set({ selectedId }),
}));

export const useSelectedBondId = () => useSelectedBondStore((state) => state.selectedId);
