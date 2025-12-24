import { LinkButton } from "@/components/reusable/link-button";
import { ClickDropdownModeToggle, HorizontalModeToggle, HoverModeToggle } from "@/components/theme_toggle";
export default function HomePage() {
    return(
        <>
        <LinkButton to="/docs" variant="destructive">
        Documentation
        </LinkButton>
        <LinkButton to="/email" variant="outline">
            Email
        </LinkButton>
        <LinkButton to="/verify">
        Verify
        </LinkButton>
        <LinkButton to="/details" variant="outline">
            User Data Input
        
        </LinkButton>
        
         <LinkButton to="/hello">
            About
        </LinkButton> 
        <LinkButton to="/login" variant="outline">
            Login
        </LinkButton>

        <LinkButton to="/verify-login">
            Login Verify
        </LinkButton>

        <LinkButton to="/forgot-password" variant="outline">
            Forgot Password
        </LinkButton>

        <LinkButton to="/app">
            App Page
        </LinkButton>

        <LinkButton to="/chat" variant="outline">
            Chat Page
        </LinkButton>

        <LinkButton to="/user">
            User Details Page
        </LinkButton>

        <HorizontalModeToggle/>
        <HoverModeToggle/>
        <ClickDropdownModeToggle/>

        </>
    )
}