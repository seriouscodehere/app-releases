import React, { useState } from 'react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { AlertCircle, User, ImageIcon } from 'lucide-react';
import { GoBackButton } from '@/components/reusable/go-back-button';

export default function UserDetailsPage() {
  const [profileImage, setProfileImage] = useState<string | null>(null);
  const [coverImage, setCoverImage] = useState<string | null>(null);
  const [username, setUsername] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [password, setPassword] = useState('');
  const [errors, setErrors] = useState<Record<string, string>>({});

  const validateUsername = (value: string): string => {
    if (!value) return 'Username is required';
    if (value.includes(' ')) return 'Username cannot contain spaces';
    if (value.includes('@')) return 'Username cannot contain @ symbol';
    if (value.length < 3) return 'Username must be at least 3 characters';
    if (value.length > 30) return 'Username must not exceed 30 characters';
    return '';
  };

  const validateFirstName = (value: string): string => {
    if (!value) return 'First name is required';
    if (value.includes(' ')) return 'First name cannot contain spaces';
    if (/\d/.test(value)) return 'First name cannot contain numbers';
    if (value.length < 3) return 'First name must be at least 3 characters';
    if (value.length > 50) return 'First name must not exceed 50 characters';
    return '';
  };

  const validateLastName = (value: string): string => {
    if (!value) return ''; // Optional field
    if (value.includes(' ')) return 'Last name cannot contain spaces';
    if (/\d/.test(value)) return 'Last name cannot contain numbers';
    if (value.length < 4) return 'Last name must be at least 4 characters';
    if (value.length > 30) return 'Last name must not exceed 30 characters';
    return '';
  };

  const validatePassword = (value: string): string => {
    if (!value) return 'Password is required';
    if (value.includes(' ')) return 'Password cannot contain spaces';
    if (value.length < 8) return 'Password must be at least 8 characters';
    if (value.length > 120) return 'Password must not exceed 120 characters';
    if (!/[a-z]/.test(value)) return 'Password must contain lowercase letters';
    if (!/[A-Z]/.test(value)) return 'Password must contain uppercase letters';
    if (!/\d/.test(value)) return 'Password must contain numbers';
    if (!/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(value)) return 'Password must contain special characters';
    return '';
  };

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>, type: 'profile' | 'cover') => {
    const file = e.target.files?.[0];
    if (!file) return;

    // Check file type
    const validTypes = ['image/jpeg', 'image/jpg', 'image/png'];
    if (!validTypes.includes(file.type)) {
      setErrors(prev => ({
        ...prev,
        [type]: 'Only JPG and PNG images are allowed'
      }));
      return;
    }

    // Clear error
    setErrors(prev => {
      const newErrors = { ...prev };
      delete newErrors[type];
      return newErrors;
    });

    // Read and display image
    const reader = new FileReader();
    reader.onload = (event) => {
      if (type === 'profile') {
        setProfileImage(event.target?.result as string);
      } else {
        setCoverImage(event.target?.result as string);
      }
    };
    reader.readAsDataURL(file);
  };

  const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Prevent @ and spaces from being typed
    if (!value.includes('@') && !value.includes(' ')) {
      setUsername(value);
      if (errors.username) {
        setErrors(prev => {
          const newErrors = { ...prev };
          delete newErrors.username;
          return newErrors;
        });
      }
    }
  };

  const handleFirstNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Prevent spaces and numbers from being typed
    if (!value.includes(' ') && !/\d/.test(value)) {
      setFirstName(value);
      if (errors.firstName) {
        setErrors(prev => {
          const newErrors = { ...prev };
          delete newErrors.firstName;
          return newErrors;
        });
      }
    }
  };

  const handleLastNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Prevent spaces and numbers from being typed
    if (!value.includes(' ') && !/\d/.test(value)) {
      setLastName(value);
      if (errors.lastName) {
        setErrors(prev => {
          const newErrors = { ...prev };
          delete newErrors.lastName;
          return newErrors;
        });
      }
    }
  };

  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    // Prevent spaces from being typed
    if (!value.includes(' ')) {
      setPassword(value);
      if (errors.password) {
        setErrors(prev => {
          const newErrors = { ...prev };
          delete newErrors.password;
          return newErrors;
        });
      }
    }
  };

  const handleSubmit = () => {
    const newErrors: Record<string, string> = {};

    // Validate images
    if (!profileImage) newErrors.profile = 'Profile image is required';
    if (!coverImage) newErrors.cover = 'Cover image is required';

    // Validate all fields
    const usernameError = validateUsername(username);
    if (usernameError) newErrors.username = usernameError;

    const firstNameError = validateFirstName(firstName);
    if (firstNameError) newErrors.firstName = firstNameError;

    const lastNameError = validateLastName(lastName);
    if (lastNameError) newErrors.lastName = lastNameError;

    const passwordError = validatePassword(password);
    if (passwordError) newErrors.password = passwordError;

    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors);
      return;
    }

    // Submit form
    console.log('Form submitted:', {
      profileImage,
      coverImage,
      username,
      firstName,
      lastName,
      password
    });

    // Success handling
    alert('Account created successfully!');
  };

  return (
    <div className="min-h-screen p-4 py-8">
      <GoBackButton/>
      <div className="max-w-2xl mx-auto">
        <Card>
          <CardHeader>
            <CardTitle className="text-3xl">Complete Your Profile</CardTitle>
            <CardDescription>
              Fill in your details to create your account
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            {/* Cover Image Upload */}
            <div className="space-y-2">
              <Label>Cover Image *</Label>
              <div className="relative h-48 border-2 border-dashed border-gray-300 rounded-lg overflow-hidden hover:border-gray-400 transition-colors">
                {coverImage ? (
                  <img src={coverImage} alt="Cover" className="w-full h-full object-cover" />
                ) : (
                  <div className="flex flex-col items-center justify-center h-full text-gray-400">
                    <ImageIcon className="h-12 w-12 mb-2" />
                    <p className="text-sm font-medium">Click to upload cover image</p>
                    <p className="text-xs">JPG or PNG only</p>
                    <p className="text-xs mt-1">Recommended: 1500 x 500 pixels</p>
                  </div>
                )}
                <input
                  type="file"
                  accept="image/jpeg,image/jpg,image/png"
                  onChange={(e) => handleImageUpload(e, 'cover')}
                  className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
                />
              </div>
              {errors.cover && (
                <p className="text-sm text-red-500 flex items-center gap-1">
                  <AlertCircle className="h-4 w-4" />
                  {errors.cover}
                </p>
              )}
            </div>

            {/* Profile Image Upload */}
            <div className="space-y-2">
              <Label>Profile Image *</Label>
              <div className="flex items-center gap-4">
                <div className="relative h-32 w-32 border-2 border-dashed border-gray-300 rounded-full overflow-hidden hover:border-gray-400 transition-colors shrink-0">
                  {profileImage ? (
                    <img src={profileImage} alt="Profile" className="w-full h-full object-cover" />
                  ) : (
                    <div className="flex flex-col items-center justify-center h-full text-gray-400">
                      <User className="h-12 w-12" />
                    </div>
                  )}
                  <input
                    type="file"
                    accept="image/jpeg,image/jpg,image/png"
                    onChange={(e) => handleImageUpload(e, 'profile')}
                    className="absolute inset-0 w-full h-full opacity-0 cursor-pointer"
                  />
                </div>
                <div className="text-sm text-gray-500">
                  <p className="font-medium">Upload your profile picture</p>
                  <p className="text-xs">JPG or PNG only</p>
                  <p className="text-xs mt-1">Recommended: 400 x 400 pixels</p>
                </div>
              </div>
              {errors.profile && (
                <p className="text-sm text-red-500 flex items-center gap-1">
                  <AlertCircle className="h-4 w-4" />
                  {errors.profile}
                </p>
              )}
            </div>

            {/* Username */}
            <div className="space-y-2">
              <Label htmlFor="username">Username *</Label>
              <Input
                id="username"
                type="text"
                placeholder="johndoe"
                value={username}
                onChange={handleUsernameChange}
                autoComplete="off"
                className={errors.username ? 'border-red-500 focus-visible:ring-red-500' : ''}
              />
              <p className="text-xs text-gray-500">
                No spaces or @ symbols allowed (3-30 characters)
              </p>
              {errors.username && (
                <p className="text-sm text-red-500 flex items-center gap-1">
                  <AlertCircle className="h-4 w-4" />
                  {errors.username}
                </p>
              )}
            </div>

            {/* First Name and Last Name */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="space-y-2">
                <Label htmlFor="firstName">First Name *</Label>
                <Input
                  id="firstName"
                  type="text"
                  placeholder="John"
                  value={firstName}
                  onChange={handleFirstNameChange}
                  autoComplete="off"
                  className={errors.firstName ? 'border-red-500 focus-visible:ring-red-500' : ''}
                />
                <p className="text-xs text-gray-500">
                  3-50 characters, no spaces or numbers
                </p>
                {errors.firstName && (
                  <p className="text-sm text-red-500 flex items-center gap-1">
                    <AlertCircle className="h-4 w-4" />
                    {errors.firstName}
                  </p>
                )}
              </div>

              <div className="space-y-2">
                <Label htmlFor="lastName">Last Name (Optional)</Label>
                <Input
                  id="lastName"
                  type="text"
                  placeholder="Doe"
                  value={lastName}
                  onChange={handleLastNameChange}
                  autoComplete="off"
                  className={errors.lastName ? 'border-red-500 focus-visible:ring-red-500' : ''}
                />
                <p className="text-xs text-gray-500">
                  4-30 characters if provided
                </p>
                {errors.lastName && (
                  <p className="text-sm text-red-500 flex items-center gap-1">
                    <AlertCircle className="h-4 w-4" />
                    {errors.lastName}
                  </p>
                )}
              </div>
            </div>

            {/* Password */}
            <div className="space-y-2">
              <Label htmlFor="password">Password *</Label>
              <Input
                id="password"
                type="password"
                placeholder="Enter your password"
                value={password}
                onChange={handlePasswordChange}
                autoComplete="off"
                className={errors.password ? 'border-red-500 focus-visible:ring-red-500' : ''}
              />
              <div className="text-xs text-gray-500 space-y-1">
                <p>Password must contain:</p>
                <ul className="list-disc list-inside ml-2 space-y-0.5">
                  <li className={password.length >= 8 && password.length <= 120 ? 'text-green-600' : ''}>
                    8-120 characters
                  </li>
                  <li className={/[a-z]/.test(password) ? 'text-green-600' : ''}>
                    Lowercase letters (a-z)
                  </li>
                  <li className={/[A-Z]/.test(password) ? 'text-green-600' : ''}>
                    Uppercase letters (A-Z)
                  </li>
                  <li className={/\d/.test(password) ? 'text-green-600' : ''}>
                    Numbers (0-9)
                  </li>
                  <li className={/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password) ? 'text-green-600' : ''}>
                    Special characters (!@#$%^&* etc.)
                  </li>
                </ul>
              </div>
              {errors.password && (
                <p className="text-sm text-red-500 flex items-center gap-1">
                  <AlertCircle className="h-4 w-4" />
                  {errors.password}
                </p>
              )}
            </div>

            {/* Submit Button */}
            <Button onClick={handleSubmit} className="w-full" size="lg">
              Create Account
            </Button>

            <div className="text-center text-sm text-gray-600">
              Already have an account?{' '}
              <a 
                href="/login" 
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = '/login';
                }}
                className="text-blue-600 hover:text-blue-700 font-medium underline cursor-pointer"
              >
                Login
              </a>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}