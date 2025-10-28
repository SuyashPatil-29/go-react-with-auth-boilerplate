import './App.css'
import type { Notebook} from './types/backend'
import { createNotebook } from './utils/notebook'
import { handleGoogleLogin, handleGoogleLogout } from './utils/auth'
import { useUser } from './hooks/auth'

function App() {

  const {user, loading} = useUser();

  const mockNotebook: Notebook = {
    id: 1,
    name: "Notebook 1",
    userId: user?.id || 1,
    chapters: [],
    createdAt: "2023-01-01T00:00:00.000Z",
    updatedAt: "2023-01-01T00:00:00.000Z",
  }

  const createNotebookCall = async () => {
    if (!user) {
      alert("Please login first")
      return
    }
    const res = await createNotebook(mockNotebook);
    console.log(res);
  }

  if (loading) {
    return <div>Loading...</div>
  }

  return (
    <>
      <div>
        {user ? (
          <div>
            <h2>Welcome, {user.name}!</h2>
            <p>Email: {user.email}</p>
            {user.imageUrl && <img src={user.imageUrl} alt={user.name} style={{ width: 50, borderRadius: '50%' }} />}
            <button onClick={handleGoogleLogout}>Logout</button>
          </div>
        ) : (
          <div>
            <h2>Please login</h2>
            <button onClick={handleGoogleLogin}>Login with Google</button>
          </div>
        )}
      </div>
      <hr />
      <button onClick={createNotebookCall}>Create notebook</button>
    </>
  )
}

export default App
