export type User = {
  id: number;
  name: string;
  email: string;
  imageUrl: string | null;
  notebooks: Notebook[];
  createdAt: string;
  updatedAt: string;
};

export type Notebook = {
  id: number;
  name: string;
  userId: number;
  chapters: Chapter[];
  createdAt: string;
  updatedAt: string;
};

export type Chapter = {
  id: number;
  name: string;
  notebookId: number;
  notebook: Notebook;
  notes: Notes[];
  createdAt: string;
  updatedAt: string;
};

export type Notes = {
  id: number;
  name: string;
  content: string;
  chapterId: number;
  chapter: Chapter;
  createdAt: string;
  updatedAt: string;
};

export type AuthenticatedUser = Pick<User, 'id' | 'name' | 'email' | 'imageUrl'>;
