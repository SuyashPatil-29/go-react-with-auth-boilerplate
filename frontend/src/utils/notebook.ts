import type { Notebook } from "@/types/backend";
import api from "@/utils/api";

export const createNotebook = async (data: Notebook) => {
    return await api.post("/notebook", data);
};

export const getNotebooks = async (id: number) => {
    return await api.delete(`/notebook/${id}`);
};

export const getNotebook = async (id: number) => {
    return await api.get(`/notebook/${id}`);
};

export const updateNotebook = async (id: number, data: Notebook) => {
    return await api.put(`/notebook/${id}`, data);
};