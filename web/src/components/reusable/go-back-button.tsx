import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"

type GoBackButtonProps = {
  className?: string
}

export function GoBackButton({ className }: GoBackButtonProps) {
  const handleGoBack = () => {
    window.history.back()
  }

  return (
    <Button
      variant="ghost"
      size="sm"
      onClick={handleGoBack}
      className={`absolute top-4 left-4  ${className ?? ""}`}
    >
      <ArrowLeft className="h-4 w-4 mr-1" />
      Go Back
    </Button>
  )
}
