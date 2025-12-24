import ForgotPasswordPage from "@/screens/auth/password/forgot_password";
import LoginPage from "@/screens/auth/signup/login";
import LoginVerifyPage from "@/screens/auth/signup/login_verify";
import UserDetailsPage from "@/screens/auth/signup/signup_steps/get_details";
import EmailInputPage from "@/screens/auth/signup/signup_steps/get_email";
import EmailVerifyPage from "@/screens/auth/signup/signup_steps/verify_email";
import AboutPage from "@/screens/landingpages/aboutPage";
import ContactPage from "@/screens/landingpages/contactPage";
import Docs from "@/screens/landingpages/docs";
import HomePage from "@/screens/landingpages/homepage";
import ChatScreen from "@/screens/main/chats/chats";
import HomeScreen from "@/screens/main/home/home_screen";
import UserProfile from "@/screens/main/profile/user_profile";

const routes = [
  { path: "/", element: <HomePage />, name: "Home" },
  { path: "/about", element: <AboutPage />, name: "About" },
  { path: "/contact", element: <ContactPage />, name: "Contact" },
  { path: "/email", element: <EmailInputPage/>, name: "Get Email" },
  { path: "/verify", element: <EmailVerifyPage/>, name: "Verify Email" },
  { path: "/details", element: <UserDetailsPage/>, name: "User Information" },
  { path: "/login", element: <LoginPage/>, name: "User Information" },
  { path: "/verify-login", element: <LoginVerifyPage/>, name: "User Information" },
  { path: "/forgot-password", element: <ForgotPasswordPage/>, name: "User Information" },


  // Docs
  { path: "/docs", element: <Docs/>, name: "Docs Screen" },

  // Main Routes

  { path: "/app", element: <HomeScreen/>, name: "Home Screen" },
  { path: "/chat", element: <ChatScreen/>, name: "Chats Screen" },
  { path: "/user", element: <UserProfile/>, name: "User Screen" },

];

export default routes;
