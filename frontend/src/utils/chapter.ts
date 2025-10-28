import api from "./api";

export const createChapter = async (data: any) => {
    return await api.post("/chapter", data);
};

export const getChapters = async (id: number) => {
    return await api.delete(`/chapter/${id}`);
};

export const getChapter = async (id: number) => {
    return await api.get(`/chapter/${id}`);
};

export const updateChapter = async (id: number, data: any) => {
    return await api.put(`/chapter/${id}`, data);
};