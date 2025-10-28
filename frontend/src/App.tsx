import type { Notebook } from '@/types/backend'
import { createNotebook } from '@/utils/notebook'
import { handleGoogleLogin, handleGoogleLogout } from '@/utils/auth'
import { useUser } from '@/hooks/auth'
import { Button } from '@/components/ui/button'
import { ModeToggle } from '@/components/ModeToggle'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { Loader, LogOut, User } from 'lucide-react'
import { useState } from 'react'
import { toast } from 'sonner'
import { Toaster } from '@/components/ui/sonner'

function App() {

  const { user, loading: userLoading } = useUser();
  const [loading, setIsLoading] = useState(false);

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
      toast.error("Please login first")
      return
    }
    try {
      setIsLoading(true);
      const res = await createNotebook(mockNotebook);
      console.log(res);
      toast.success("Notebook created successfully!");
    } catch (error: any) {
      // Handle axios/HTTP errors
      if (error.response) {
        // Server responded with error status
        const status = error.response.status;
        const message = error.response.data?.message || error.response.data?.error || error.message;
        toast.error(`Error ${status}: ${message}`);
      } else if (error.request) {
        // Request made but no response
        toast.error("No response from server. Please check your connection.");
      } else {
        // Something else happened
        toast.error(error.message || "An unexpected error occurred");
      }
      console.error("Error creating notebook:", error);
    } finally {
      setIsLoading(false);
    }
  }

  if (userLoading) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <div className="text-muted-foreground text-lg">Loading...</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-background">
      <Toaster />
      {/* Header */}
      <header className="border-b border-border">
        <div className="max-w-6xl mx-auto px-6 py-4 flex items-center justify-between">
          <h1 className="text-2xl font-bold text-foreground">Notes App</h1>

          <div className="flex items-center gap-3">
            <ModeToggle />
            {user && (
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <button className="flex items-center gap-2 rounded-full hover:opacity-80 transition-opacity focus:outline-none focus:ring-2 focus:ring-ring">
                    {user.imageUrl ? (
                      <img
                        src={user.imageUrl}
                        alt={user.name}
                        className="w-9 h-9 rounded-full ring-2 ring-border"
                      />
                    ) : (
                      <div className="w-9 h-9 rounded-full bg-primary flex items-center justify-center ring-2 ring-border">
                        <User className="w-5 h-5 text-primary-foreground" />
                      </div>
                    )}
                  </button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="start" className="w-56">
                  <DropdownMenuLabel className="font-normal">
                    <div className="flex flex-col space-y-1">
                      <p className="text-sm font-medium leading-none">{user.name}</p>
                      <p className="text-xs leading-none text-muted-foreground">
                        {user.email}
                      </p>
                    </div>
                  </DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem onClick={handleGoogleLogout} className="cursor-pointer">
                    <LogOut className="mr-2 h-4 w-4" />
                    <span>Log out</span>
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            )}
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-6xl mx-auto px-6 py-12">
        {user ? (
          <div className="space-y-8">
            {/* Welcome Section */}
            <div className="space-y-2">
              <h2 className="text-3xl font-bold text-foreground">
                Welcome back, {user.name}!
              </h2>
              <p className="text-muted-foreground">{user.email}</p>
            </div>

            {/* Action Cards */}
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
              <div className="bg-card border border-border rounded-lg p-6 space-y-3 hover:border-primary/50 transition-colors">
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold text-card-foreground">Create Notebook</h3>
                  <p className="text-sm text-muted-foreground">
                    Start organizing your notes in a new notebook
                  </p>
                </div>
                <Button
                  onClick={createNotebookCall}
                  disabled={loading}
                  className="w-full flex items-center justify-center"
                >
                  {loading && <Loader className="mr-2 h-4 w-4 animate-spin" size={20} />}
                  Create Notebook
                </Button>
              </div>

              <div className="bg-card border border-border rounded-lg p-6 space-y-3">
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold text-card-foreground">Recent Notes</h3>
                  <p className="text-sm text-muted-foreground">
                    Access your recently edited notes
                  </p>
                </div>
                <Button variant="outline" className="w-full" disabled>
                  Coming Soon
                </Button>
              </div>

              <div className="bg-card border border-border rounded-lg p-6 space-y-3">
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold text-card-foreground">Search</h3>
                  <p className="text-sm text-muted-foreground">
                    Find notes across all your notebooks
                  </p>
                </div>
                <Button variant="outline" className="w-full" disabled>
                  Coming Soon
                </Button>
              </div>
            </div>
          </div>
        ) : (
          <div className="flex flex-col items-center justify-center min-h-[60vh] space-y-6">
            <div className="text-center space-y-3">
              <h2 className="text-4xl font-bold text-foreground">Welcome to Notes App</h2>
              <p className="text-lg text-muted-foreground max-w-md">
                Sign in with your Google account to start organizing your thoughts and ideas
              </p>
            </div>
            <Button size="lg" onClick={handleGoogleLogin}>
              Login with Google
            </Button>
          </div>
        )}
      </main>
    </div>
  )
}

export default App
