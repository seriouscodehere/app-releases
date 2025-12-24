import { useState, useRef, useEffect, type JSX } from "react"
import { Moon, Sun, Monitor, ChevronDown } from "lucide-react"
import { Button } from "@/components/ui/button"
import { useTheme } from "@/components/theme_provider"

type Theme = "light" | "dark" | "system"

const modes: { name: Theme; icon: JSX.Element; label: string }[] = [
  { name: "light", icon: <Sun className="h-[1.1rem] w-[1.1rem]" />, label: "Light" },
  { name: "dark", icon: <Moon className="h-[1.1rem] w-[1.1rem]" />, label: "Dark" },
  { name: "system", icon: <Monitor className="h-[1.1rem] w-[1.1rem]" />, label: "System" },
]

// 1. Horizontal simple toggle
export function HorizontalModeToggle() {
  const { theme, setTheme } = useTheme()

  return (
    <div className="flex items-center border rounded-3xl p-1 gap-1 max-w-26">
      {modes.map((mode) => (
        <Button
          key={mode.name}
          variant="ghost"
          size="icon"
          onClick={() => setTheme(mode.name)}
          className={`h-7 w-7 ${theme === mode.name ? "bg-accent" : ""}`}
        >
          {mode.icon}
          <span className="sr-only">{mode.label}</span>
        </Button>
      ))}
    </div>
  )
}

// 2. Click dropdown toggle (close on outside click)
export function ClickDropdownModeToggle() {
  const { theme, setTheme } = useTheme()
  const [showMenu, setShowMenu] = useState(false)
  const ref = useRef<HTMLDivElement>(null)

  useEffect(() => {
    function handleClickOutside(event: MouseEvent) {
      if (ref.current && !ref.current.contains(event.target as Node)) {
        setShowMenu(false)
      }
    }
    document.addEventListener("mousedown", handleClickOutside)
    return () => document.removeEventListener("mousedown", handleClickOutside)
  }, [])

  return (
    <div className="relative inline-block" ref={ref}>
      <Button onClick={() => setShowMenu(!showMenu)} variant="ghost">
        {modes.find((m) => m.name === theme)?.icon}
      </Button>
      {showMenu && (
        <div className="absolute mt-2 right-0 bg-background border border-gray-300 rounded-lg shadow-lg min-w-35 z-50">
          {modes.map((mode) => (
            <Button
              key={mode.name}
              variant="ghost"
              className={`w-full justify-start px-3 py-2 rounded hover:bg-accent ${theme === mode.name ? "bg-accent" : ""}`}
              onClick={() => {
                setTheme(mode.name)
                setShowMenu(false)
              }}
            >
              {mode.icon}
              <span className="ml-2">{mode.label}</span>
            </Button>
          ))}
        </div>
      )}
    </div>
  )
}

// 3. Hover menu toggle (side menu, touching the button)
export function HoverModeToggle() {
  const { theme, setTheme } = useTheme()
  const [showMenu, setShowMenu] = useState(false)
  const containerRef = useRef<HTMLDivElement>(null)

  return (
    <div
      className="relative inline-block"
      ref={containerRef}
      onMouseEnter={() => setShowMenu(true)}
      onMouseLeave={() => setShowMenu(false)}
    >
      <Button className="flex items-center gap-2 rounded-full border px-3 py-1">
        <span>{modes.find((m) => m.name === theme)?.label}</span>
        <ChevronDown className={`w-4 h-4 transition-transform duration-200 ${showMenu ? "rotate-180" : ""}`} />
      </Button>
      {showMenu && (
        <div className="absolute top-0 left-full bg-background border border-gray-300 rounded-lg shadow-lg min-w-30 z-50">
          {modes.map((mode) => (
            <Button
              key={mode.name}
              variant="ghost"
              className={`w-full justify-start px-3 py-2 rounded hover:bg-accent ${theme === mode.name ? "bg-accent" : ""}`}
              onClick={() => setTheme(mode.name)}
            >
              {mode.label}
            </Button>
          ))}
        </div>
      )}
    </div>
  )
}
