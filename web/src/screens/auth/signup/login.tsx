import React, { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { AlertCircle, CheckCircle2, Eye, EyeOff } from 'lucide-react';
import { LinkButton } from '@/components/reusable/link-button';
import { GoBackButton } from '@/components/reusable/go-back-button';

export default function LoginPage() {
  const [emailOrUsername, setEmailOrUsername] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [errors, setErrors] = useState({ emailOrUsername: '', password: '' });
  const [submitted, setSubmitted] = useState(false);

  const validateEmailOrUsername = (value: string): string => {
    if (!value) {
      return 'Email or username is required';
    }

    if (value.length < 3) {
      return 'Must be at least 3 characters long';
    }

    if (value.length > 100) {
      return 'Must not exceed 100 characters';
    }

    return '';
  };

  const validatePassword = (value: string): string => {
    if (!value) {
      return 'Password is required';
    }

    if (value.length < 8 || value.length > 60) {
      return 'Password must be between 8 and 60 characters';
    }

    if (!/[A-Z]/.test(value)) {
      return 'Password must contain at least one uppercase letter';
    }

    if (!/[a-z]/.test(value)) {
      return 'Password must contain at least one lowercase letter';
    }

    if (!/[0-9]/.test(value)) {
      return 'Password must contain at least one number';
    }

    if (!/[!@#$%^&*(),.?":{}|<>]/.test(value)) {
      return 'Password must contain at least one special character';
    }

    return '';
  };

  const handleEmailOrUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setEmailOrUsername(value);
    
    if (errors.emailOrUsername) {
      setErrors(prev => ({ ...prev, emailOrUsername: '' }));
    }
    
    if (submitted) {
      setSubmitted(false);
    }
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setPassword(value);
    
    if (errors.password) {
      setErrors(prev => ({ ...prev, password: '' }));
    }
    
    if (submitted) {
      setSubmitted(false);
    }
  };

  const handleSubmit = () => {
    const emailOrUsernameError = validateEmailOrUsername(emailOrUsername);
    const passwordError = validatePassword(password);
    
    if (emailOrUsernameError || passwordError) {
      setErrors({
        emailOrUsername: emailOrUsernameError,
        password: passwordError
      });
      return;
    }

    // Valid credentials
    setErrors({ emailOrUsername: '', password: '' });
    setSubmitted(true);
    console.log('Login submitted:', { emailOrUsername, password });
    
    // Reset after 3 seconds
    setTimeout(() => {
      setSubmitted(false);
      setEmailOrUsername('');
      setPassword('');
    }, 3000);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleSubmit();
    }
  };

  return (
    <div className="relative min-h-screen p-4">
      <GoBackButton/>
      <div className="flex min-h-screen items-center justify-center"> 
        <Card className="w-full max-w-md">
          <CardHeader>
            <CardTitle className="text-2xl">Recall Your Memory</CardTitle>
            <CardDescription>
              Solve Problem Below To Get Access
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="emailOrUsername">Email or Username</Label>
                <Input
                  id="emailOrUsername"
                  type="text"
                  placeholder="you@example.com or username"
                  value={emailOrUsername}
                  onChange={handleEmailOrUsernameChange}
                  onKeyPress={handleKeyPress}
                  autoComplete="username"
                  className={`w-full ${errors.emailOrUsername ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
                />
                <p className="text-xs text-gray-500">
                  Must be 3-100 characters
                </p>
              </div>

              <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <div className="relative">
                  <Input
                    id="password"
                    type={showPassword ? 'text' : 'password'}
                    placeholder="Enter your password"
                    value={password}
                    onChange={handlePasswordChange}
                    onKeyPress={handleKeyPress}
                    autoComplete="current-password"
                    className={`w-full pr-10 ${errors.password ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
                  />
                  <button
                    type="button"
                    onClick={() => setShowPassword(!showPassword)}
                    className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                  >
                    {showPassword ? (
                      <EyeOff className="h-4 w-4" />
                    ) : (
                      <Eye className="h-4 w-4" />
                    )}
                  </button>
                </div>
                <p className="text-xs text-gray-500">
                  8-60 characters with uppercase, lowercase, number & special character
                </p>
              </div>

              {(errors.emailOrUsername || errors.password) && (
                <Alert variant="destructive" className="flex items-start gap-3">
                  <AlertCircle className="h-5 w-5 shrink-0 mt-0.5" />
                  <div className="space-y-1">
                    {errors.emailOrUsername && (
                      <AlertDescription className="text-sm">{errors.emailOrUsername}</AlertDescription>
                    )}
                    {errors.password && (
                      <AlertDescription className="text-sm">{errors.password}</AlertDescription>
                    )}
                  </div>
                </Alert>
              )}

              {submitted && (
                <Alert className="border-green-500 bg-green-50 flex items-center gap-3">
                  <CheckCircle2 className="h-8 w-8 text-green-600 shrink-0" />
                  <AlertDescription className="text-green-600 text-base">
                    Login successful! Redirecting...
                  </AlertDescription>
                </Alert>
              )}

              <Button 
                onClick={handleSubmit} 
                className="w-full"
                disabled={submitted}
              >
                {submitted ? 'Logging in...' : 'Login'}
              </Button>
              <div className="text-center">
                <LinkButton to="/email" variant="link">Create Account</LinkButton>
                <LinkButton to="/forgot-password" variant="link">Forgot Password</LinkButton>
                <LinkButton to="" variant="link">Privacy Policies</LinkButton>
              </div>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}