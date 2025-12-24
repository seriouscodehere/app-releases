import React, { useState, useEffect } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { AlertCircle, CheckCircle2, Mail, Eye, EyeOff, Check, X } from 'lucide-react';
import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from '@/components/ui/input-otp';
import { GoBackButton } from '@/components/reusable/go-back-button';
import { LinkButton } from '@/components/reusable/link-button';

export default function ForgotPasswordPage() {
  const [email, setEmail] = useState('');
  const [otp, setOtp] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [errors, setErrors] = useState({ email: '', otp: '', newPassword: '', confirmPassword: '' });
  const [otpSent, setOtpSent] = useState(false);
  const [otpVerified, setOtpVerified] = useState(false);
  const [passwordReset, setPasswordReset] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [passwordValidation, setPasswordValidation] = useState({
    length: false,
    uppercase: false,
    lowercase: false,
    number: false,
    special: false
  });

  useEffect(() => {
    if (newPassword) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setPasswordValidation({
        length: newPassword.length >= 8 && newPassword.length <= 60,
        uppercase: /[A-Z]/.test(newPassword),
        lowercase: /[a-z]/.test(newPassword),
        number: /[0-9]/.test(newPassword),
        special: /[!@#$%^&*(),.?":{}|<>]/.test(newPassword)
      });
    }
  }, [newPassword]);

  const validateEmail = (value: string): string => {
    if (!value) {
      return 'Email is required';
    }

    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(value)) {
      return 'Please enter a valid email address';
    }

    if (value.length > 100) {
      return 'Email must not exceed 100 characters';
    }

    return '';
  };

  const validateOtp = (value: string): string => {
    if (!value) {
      return 'OTP is required';
    }

    if (value.length !== 6) {
      return 'OTP must be exactly 6 digits';
    }

    return '';
  };

  const validateNewPassword = (value: string): string => {
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

  const validateConfirmPassword = (value: string, password: string): string => {
    if (!value) {
      return 'Please confirm your password';
    }

    if (value !== password) {
      return 'Passwords do not match';
    }

    return '';
  };

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setEmail(value);
    
    if (errors.email) {
      setErrors(prev => ({ ...prev, email: '' }));
    }
  };

  const handleOtpChange = (value: string) => {
    setOtp(value);
    
    if (errors.otp) {
      setErrors(prev => ({ ...prev, otp: '' }));
    }
  };

  const handleNewPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setNewPassword(value);
    
    if (errors.newPassword) {
      setErrors(prev => ({ ...prev, newPassword: '' }));
    }
  };

  const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setConfirmPassword(value);
    
    if (errors.confirmPassword) {
      setErrors(prev => ({ ...prev, confirmPassword: '' }));
    }
  };

  const handleSendOtp = () => {
    const emailError = validateEmail(email);
    
    if (emailError) {
      setErrors({ email: emailError, otp: '', newPassword: '', confirmPassword: '' });
      return;
    }

    setIsLoading(true);
    setErrors({ email: '', otp: '', newPassword: '', confirmPassword: '' });
    
    // Simulate API call
    setTimeout(() => {
      setOtpSent(true);
      setIsLoading(false);
      console.log('OTP sent to:', email);
    }, 1500);
  };

  const handleVerifyOtp = () => {
    const otpError = validateOtp(otp);
    
    if (otpError) {
      setErrors({ email: '', otp: otpError, newPassword: '', confirmPassword: '' });
      return;
    }

    setIsLoading(true);
    setErrors({ email: '', otp: '', newPassword: '', confirmPassword: '' });
    
    // Simulate API call
    setTimeout(() => {
      setOtpVerified(true);
      setIsLoading(false);
      console.log('OTP verified:', otp);
    }, 1500);
  };

  const handleResetPassword = () => {
    const newPasswordError = validateNewPassword(newPassword);
    const confirmPasswordError = validateConfirmPassword(confirmPassword, newPassword);
    
    if (newPasswordError || confirmPasswordError) {
      setErrors({ 
        email: '', 
        otp: '', 
        newPassword: newPasswordError, 
        confirmPassword: confirmPasswordError 
      });
      return;
    }

    setIsLoading(true);
    setErrors({ email: '', otp: '', newPassword: '', confirmPassword: '' });
    
    // Simulate API call
    setTimeout(() => {
      setPasswordReset(true);
      setIsLoading(false);
      console.log('Password reset successful');
      
      // Reset after 3 seconds
      setTimeout(() => {
        setEmail('');
        setOtp('');
        setNewPassword('');
        setConfirmPassword('');
        setOtpSent(false);
        setOtpVerified(false);
        setPasswordReset(false);
      }, 3000);
    }, 1500);
  };

  const handleResendCode = () => {
    setOtp('');
    setErrors({ email: '', otp: '', newPassword: '', confirmPassword: '' });
    handleSendOtp();
  };

  return (
    <div className="flex min-h-screen items-center justify-center p-4">
      <GoBackButton/>
      <Card className="w-full max-w-md">
        <CardHeader>
          <CardTitle className="text-2xl">Forgot Password?</CardTitle>
          <CardDescription>
            {!otpSent ? 'Enter your email to receive a reset code' : 
             !otpVerified ? 'Enter the verification code' : 
             'Create your new password'}
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="email">Email Address</Label>
              <div className="relative">
                <Mail className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-gray-500" />
                <Input
                  id="email"
                  type="email"
                  placeholder="you@example.com"
                  value={email}
                  onChange={handleEmailChange}
                  autoComplete="email"
                  disabled={otpSent}
                  className={`pl-10 ${errors.email ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
                />
              </div>
              {!otpSent && (
                <p className="text-xs text-gray-500">
                  We'll send a 6-digit verification code to this email
                </p>
              )}
              {otpSent && (
                <p className="text-xs text-green-600">
                  ✓ Code sent to {email}
                </p>
              )}
            </div>

            {errors.email && (
              <Alert variant="destructive" className="flex items-start gap-3">
                <AlertCircle className="h-5 w-5 shrink-0 mt-0.5" />
                <AlertDescription className="text-sm">{errors.email}</AlertDescription>
              </Alert>
            )}

            {!otpSent ? (
              <Button 
                onClick={handleSendOtp} 
                className="w-full"
                disabled={isLoading}
              >
                {isLoading ? 'Sending...' : 'Send Reset Code'}
              </Button>
            ) : !otpVerified ? (
              <>
                <div className="space-y-2">
                  <Label htmlFor="otp">Verification Code</Label>
                  <div className="flex justify-center">
                    <InputOTP
                      maxLength={6}
                      value={otp}
                      onChange={handleOtpChange}
                      onPaste={(e) => e.preventDefault()}
                    >
                      <InputOTPGroup>
                        <InputOTPSlot index={0} />
                        <InputOTPSlot index={1} />
                        <InputOTPSlot index={2} />
                        <InputOTPSlot index={3} />
                        <InputOTPSlot index={4} />
                        <InputOTPSlot index={5} />
                      </InputOTPGroup>
                    </InputOTP>
                  </div>
                  <p className="text-xs text-gray-500 text-center">
                    Enter the 6-digit code (typing only, paste disabled)
                  </p>
                </div>

                {errors.otp && (
                  <Alert variant="destructive" className="flex items-start gap-3">
                    <AlertCircle className="h-5 w-5 shrink-0 mt-0.5" />
                    <AlertDescription className="text-sm">{errors.otp}</AlertDescription>
                  </Alert>
                )}

                <Button 
                  onClick={handleVerifyOtp} 
                  className="w-full"
                  disabled={isLoading || otp.length !== 6}
                >
                  {isLoading ? 'Verifying...' : 'Verify Code'}
                </Button>

                <div className="text-center">
                  <Button 
                    variant="link" 
                    onClick={() => {
                      setOtpSent(false);
                      setOtp('');
                      setErrors({ email: '', otp: '', newPassword: '', confirmPassword: '' });
                    }}
                    className="text-sm"
                  >
                    Change email
                  </Button>
                  <span className="text-gray-400 mx-2">•</span>
                  <Button 
                    variant="link" 
                    onClick={handleResendCode}
                    className="text-sm"
                    disabled={isLoading}
                  >
                    Resend code
                  </Button>
                </div>
              </>
            ) : (
              <>
                <div className="space-y-2">
                  <Label htmlFor="newPassword">New Password</Label>
                  <div className="relative">
                    <Input
                      id="newPassword"
                      type={showNewPassword ? 'text' : 'password'}
                      placeholder="Enter new password"
                      value={newPassword}
                      onChange={handleNewPasswordChange}
                      autoComplete="new-password"
                      className={`pr-10 ${errors.newPassword ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
                    />
                    <button
                      type="button"
                      onClick={() => setShowNewPassword(!showNewPassword)}
                      className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                    >
                      {showNewPassword ? (
                        <EyeOff className="h-4 w-4" />
                      ) : (
                        <Eye className="h-4 w-4" />
                      )}
                    </button>
                  </div>
                </div>

                <div className="space-y-1 text-xs">
                  <div className={`flex items-center gap-2 ${passwordValidation.length ? 'text-green-600' : 'text-gray-500'}`}>
                    {passwordValidation.length ? <Check className="h-3 w-3" /> : <X className="h-3 w-3" />}
                    <span>8-60 characters</span>
                  </div>
                  <div className={`flex items-center gap-2 ${passwordValidation.uppercase ? 'text-green-600' : 'text-gray-500'}`}>
                    {passwordValidation.uppercase ? <Check className="h-3 w-3" /> : <X className="h-3 w-3" />}
                    <span>At least one uppercase letter (A-Z)</span>
                  </div>
                  <div className={`flex items-center gap-2 ${passwordValidation.lowercase ? 'text-green-600' : 'text-gray-500'}`}>
                    {passwordValidation.lowercase ? <Check className="h-3 w-3" /> : <X className="h-3 w-3" />}
                    <span>At least one lowercase letter (a-z)</span>
                  </div>
                  <div className={`flex items-center gap-2 ${passwordValidation.number ? 'text-green-600' : 'text-gray-500'}`}>
                    {passwordValidation.number ? <Check className="h-3 w-3" /> : <X className="h-3 w-3" />}
                    <span>At least one number (0-9)</span>
                  </div>
                  <div className={`flex items-center gap-2 ${passwordValidation.special ? 'text-green-600' : 'text-gray-500'}`}>
                    {passwordValidation.special ? <Check className="h-3 w-3" /> : <X className="h-3 w-3" />}
                    <span>At least one special character (!@#$%^&*)</span>
                  </div>
                </div>

                <div className="space-y-2">
                  <Label htmlFor="confirmPassword">Confirm Password</Label>
                  <div className="relative">
                    <Input
                      id="confirmPassword"
                      type={showConfirmPassword ? 'text' : 'password'}
                      placeholder="Confirm new password"
                      value={confirmPassword}
                      onChange={handleConfirmPasswordChange}
                      autoComplete="new-password"
                      className={`pr-10 ${errors.confirmPassword ? 'border-red-500 focus-visible:ring-red-500' : ''}`}
                    />
                    <button
                      type="button"
                      onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                      className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-700"
                    >
                      {showConfirmPassword ? (
                        <EyeOff className="h-4 w-4" />
                      ) : (
                        <Eye className="h-4 w-4" />
                      )}
                    </button>
                  </div>
                </div>

                {(errors.newPassword || errors.confirmPassword) && (
                  <Alert variant="destructive" className="flex items-start gap-3">
                    <AlertCircle className="h-5 w-5 shrink-0 mt-0.5" />
                    <div className="space-y-1">
                      {errors.newPassword && (
                        <AlertDescription className="text-sm">{errors.newPassword}</AlertDescription>
                      )}
                      {errors.confirmPassword && (
                        <AlertDescription className="text-sm">{errors.confirmPassword}</AlertDescription>
                      )}
                    </div>
                  </Alert>
                )}

                {passwordReset && (
                  <Alert className="border-green-500 bg-green-50 flex items-center gap-3">
                    <CheckCircle2 className="h-8 w-8 text-green-600 shrink-0" />
                    <AlertDescription className="text-green-600 text-base">
                      Password reset successful! Redirecting to login...
                    </AlertDescription>
                  </Alert>
                )}

                <Button 
                  onClick={handleResetPassword} 
                  className="w-full"
                  disabled={isLoading || passwordReset}
                >
                  {isLoading ? 'Resetting...' : passwordReset ? 'Success!' : 'Reset Password'}
                </Button>
              </>
            )}

            <div className="text-center pt-2">
              <LinkButton to="/login" variant="link">
              Back to Login
              </LinkButton>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}