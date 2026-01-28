import { create } from 'zustand'

interface AuthState {
    userId: string | null;
    token: string | null;
    setUserId: (id: string | null) => void;
    setToken: (token: string | null) => void;
}

export const useAuthStore = create<AuthState>((set) => ({
    userId: null,
    token: null,
    setUserId: (id) => set({ userId: id }),
    setToken: (token) => set({ token: token }),
}));