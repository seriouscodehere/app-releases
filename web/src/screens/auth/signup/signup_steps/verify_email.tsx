import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Alert, AlertDescription } from '@/components/ui/alert';
import { InputOTP, InputOTPGroup, InputOTPSlot } from '@/components/ui/input-otp';
import { AlertCircle, CheckCircle2, Mail } from 'lucide-react';
import { GoBackButton } from '@/components/reusable/go-back-button';

export default function EmailVerifyPage() {
  const [otp, setOtp] = useState('');
  const [error, setError] = useState('');
  const [verified, setVerified] = useState(false);
  const [resending, setResending] = useState(false);

  const validateOTP = (otp: string): string => {
    // Check if OTP is empty
    if (!otp) {
      return 'OTP is required';
    }

    // Check if OTP is exactly 6 digits
    if (otp.length !== 6) {
      return 'OTP must be exactly 6 digits';
    }

    // Check if OTP contains only numbers
    if (!/^\d+$/.test(otp)) {
      return 'OTP must contain only numbers';
    }

    return '';
  };

  const handleOtpChange = (value: string) => {
    setOtp(value);
    
    // Clear error when user starts typing
    if (error) {
      setError('');
    }
    
    // Clear verified state when user modifies OTP
    if (verified) {
      setVerified(false);
    }
  };

  const handlePaste = (e: React.ClipboardEvent) => {
    e.preventDefault();
    return false;
  };

  const handleVerify = () => {
    const validationError = validateOTP(otp);
    
    if (validationError) {
      setError(validationError);
      return;
    }

    // OTP is valid
    setError('');
    setVerified(true);
    console.log('OTP verified:', otp);
    
    // Redirect or perform action after 2 seconds
    setTimeout(() => {
      console.log('Redirecting to dashboard...');
      // window.location.href = '/dashboard';
    }, 2000);
  };

  const handleResendOTP = () => {
    setResending(true);
    console.log('Resending OTP...');
    
    // Simulate resend delay
    setTimeout(() => {
      setResending(false);
      setOtp('');
      setError('');
      setVerified(false);
    }, 2000);
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <GoBackButton/>
      <Card className="w-full max-w-md">
        <CardHeader>
          <div className="flex items-center gap-2">
            <Mail className="h-6 w-6 text-purple-600" />
            <CardTitle className="text-2xl">Verify Your Email</CardTitle>
          </div>
          <CardDescription>
            We've sent a 6-digit verification code to your email address
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="otp" className="text-center block">Enter OTP Code</Label>
              <div className="flex justify-center">
                <InputOTP
                  maxLength={6}
                  value={otp}
                  onChange={handleOtpChange}
                  onPaste={handlePaste}
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
                Enter the 6-digit code sent to your email
              </p>
            </div>

            {error && (
              <Alert variant="destructive" className="flex items-center gap-3">
                <AlertCircle className="h-8 w-8 shrink-0" />
                <AlertDescription className="text-base">{error}</AlertDescription>
              </Alert>
            )}

            {verified && (
              <Alert className="border-green-500 bg-green-50 flex items-center gap-3">
                <CheckCircle2 className="h-8 w-8 text-green-600 shrink-0" />
                <AlertDescription className="text-green-600 text-base">
                  Email verified successfully! Redirecting...
                </AlertDescription>
              </Alert>
            )}

          <Button 
            onClick={handleVerify} 
            className="flex items-center justify-center mx-auto w-40"
            disabled={verified}
          >
           {verified ? 'Verified!' : 'Verify Email'}
          </Button>
            <div className="text-center">
              <button
                onClick={handleResendOTP}
                disabled={resending || verified}
                className="text-sm text-purple-600 hover:text-purple-700 font-medium underline disabled:text-gray-400 disabled:no-underline"
              >
                {resending ? 'Resending...' : "Didn't receive code? Resend"}
              </button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}