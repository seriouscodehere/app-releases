import { GoBackButton } from "@/components/reusable/go-back-button";
import { HorizontalModeToggle } from "@/components/theme_toggle";

export default function Docs() {
  return (
  <>
    <div className="p-8 min-h-screen">
  <div className="mb-10">
  <GoBackButton/>
  </div>
      <h1 className="text-3xl font-bold mb-6">Authentication Pages Documentation</h1>

      {/* Introduction */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Introduction</h2>
        <p>
          This documentation covers the authentication pages of our application, including Login, Register, Forgot Password, and User Profile.
          Each section describes the page structure, props, and behavior.
        </p>
      </section>

      {/* Login Page */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Login Page</h2>
        <p>
          The Login page allows users to access the app using their credentials.
        </p>
        <ul className="list-disc list-inside ml-4">
          <li>Fields: Email, Password</li>
          <li>Buttons: Login, Forgot Password</li>
          <li>Validation: Required fields, valid email format</li>
        </ul>
      </section>

      {/* Register Page */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Register Page</h2>
        <p>
          The Register page allows new users to create an account.
        </p>
        <ul className="list-disc list-inside ml-4">
          <li>Fields: Name, Email, Password, Confirm Password</li>
          <li>Buttons: Register, Login link</li>
          <li>Validation: Required fields, password match, valid email</li>
        </ul>
      </section>

      {/* Forgot Password Page */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Forgot Password Page</h2>
        <p>
          This page allows users to reset their password via email.
        </p>
        <ul className="list-disc list-inside ml-4">
          <li>Fields: Email</li>
          <li>Buttons: Send Reset Link</li>
          <li>Validation: Required email, valid email format</li>
        </ul>
      </section>

      {/* Profile Page */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Profile Page</h2>
        <p>
          The Profile page displays user information and allows updates.
        </p>
        <ul className="list-disc list-inside ml-4">
          <li>Fields: Name, Email, Profile Picture</li>
          <li>Buttons: Update Profile, Change Password</li>
          <li>Validation: Required fields, valid email</li>
        </ul>
      </section>

      {/* Notes */}
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-2">Notes</h2>
        <p>
          All authentication forms should handle loading states and error messages properly. Use client-side validation for better UX, and server-side validation for security.
        </p>
      </section>
    </div>
    <div className="p-2 bg-amber-100 m-4 border rounded-3xl text-gray-900">

    <p className="text-red-500">!Important Note:</p>
    <p>The authentication Pages are not linked to Backend Yet but backend struture for authentication is fully matured and tested</p>
    </div>
    <HorizontalModeToggle/>
    </>
  )
}
