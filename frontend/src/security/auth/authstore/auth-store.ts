import { create } from 'zustand'

interface AuthState {
    userId: string | null;
    token: string | null;
    setUserId: (id: string | undefined) => void;
    setToken: (token: string | undefined) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
    userId: null,
    token: null,
    setUserId: (id) => set({ userId: id }),
    setToken: (token) => set({ token }),
}));