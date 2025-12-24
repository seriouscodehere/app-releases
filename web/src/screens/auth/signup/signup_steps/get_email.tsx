import React, { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { AlertCircle, CheckCircle2 } from 'lucide-react';
import { GoBackButton } from '@/components/reusable/go-back-button';
import { LinkButton } from '@/components/reusable/link-button';

export default function EmailInputPage() {
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const [submitted, setSubmitted] = useState(false);

  const validateEmail = (email: string): string => {
    // Check if email is empty
    if (!email) {
      return 'Email is required';
    }

    // Check minimum length
    if (email.length < 8) {
      return 'Email must be at least 8 characters long';
    }

    // Check maximum length
    if (email.length > 100) {
      return 'Email must not exceed 100 characters';
    }

    // Check valid email format
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      return 'Please enter a valid email address';
    }

    return '';
  };

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setEmail(value);
    
    // Clear error when user starts typing
    if (error) {
      setError('');
    }
    
    // Clear submitted state when user modifies email
    if (submitted) {
      setSubmitted(false);
    }
  };

  const handleSubmit = () => {
    const validationError = validateEmail(email);
    
    if (validationError) {
      setError(validationError);
      return;
    }

    // Email is valid
    setError('');
    setSubmitted(true);
    console.log('Email submitted:', email);
    
    // Reset after 3 seconds
    setTimeout(() => {
      setSubmitted(false);
      setEmail('');
    }, 3000);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleSubmit();
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <GoBackButton/>
      <Card className="w-full max-w-md">
        <CardHeader>

          <CardTitle className="text-2xl">Stay Connected</CardTitle>
          <CardDescription>
            Enter your email to receive updates and notifications
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="email">Email Address</Label>
              <Input
                id="email"
                type="email"
                placeholder="you@example.com"
                value={email}
                onChange={handleEmailChange}
                onKeyPress={handleKeyPress}
                autoComplete="off"
                className={`w-full ${error ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
              />
              <p className="text-xs text-gray-500">
                Must be 8-100 characters and a valid email format
              </p>
            </div>

            {error && (
              <Alert variant="destructive" className="flex items-center gap-3">
                <AlertCircle className="h-8 w-8 shrink-0" />
                <AlertDescription className="text-base">{error}</AlertDescription>
              </Alert>
            )}

            {submitted && (
              <Alert className="border-green-500 bg-green-50 flex items-center gap-3">
                <CheckCircle2 className="h-8 w-8 text-green-600 shrink-0" />
                <AlertDescription className="text-green-600 text-base">
                  Thank you! We've received your email.
                </AlertDescription>
              </Alert>
            )}

            <Button 
              onClick={handleSubmit} 
              className="w-full"
              disabled={submitted}
            >
              {submitted ? 'Submitted!' : 'Subscribe'}
            </Button>

            <div className="text-center text-sm text-gray-600">
              Already have an account?{' '}
            <LinkButton to="/login" variant="link">
            Login
            </LinkButton>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}