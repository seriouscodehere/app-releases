import { ThemeProvider } from "@/components/theme_provider"
import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import routes from "./routing/routes"
import NotFoundPage from "./screens/notFound/notFound"

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <Router>
        <Routes>
          {routes.map((r) => (
            <Route key={r.path} path={r.path} element={r.element} />
          ))}
           <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </Router>
    </ThemeProvider>
  )
}

export default App
