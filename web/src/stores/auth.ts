import { create } from 'zustand';

type AuthState = {
  isUserLoggedIn: boolean;
  setIsUserLoggedIn: (value: boolean) => void;
};

export const useAuthStore = create<AuthState>((set) => ({
  isUserLoggedIn: false,
  setIsUserLoggedIn: (isUserLoggedIn) => set({ isUserLoggedIn }),
}));

export const useIsUserLoggedIn = () => useAuthStore((state) => state.isUserLoggedIn);
