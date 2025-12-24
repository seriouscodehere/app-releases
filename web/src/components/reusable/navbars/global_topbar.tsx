import { useState } from "react";
import {
  Bell,
  MessageSquare,
  User,
  Settings,
  LogOut,
  CreditCard,
  Users,
  BarChart,
  Shield,
  Mail,
  Calendar,
  Download,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Separator } from "@/components/ui/separator";
import { ClickDropdownModeToggle, HorizontalModeToggle } from "@/components/theme_toggle";
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from "@/components/ui/sheet";
import { Link } from "react-router-dom";

interface Notification {
  id: number;
  title: string;
  description: string;
  time: string;
  read: boolean;
}

export function GlobalTopbar() {
  const [notifications] = useState<Notification[]>([
    { id: 1, title: "New Message", description: "You have a new message from John", time: "5 min ago", read: false },
    { id: 2, title: "System Update", description: "System maintenance scheduled", time: "1 hour ago", read: true },
    { id: 3, title: "Payment Received", description: "Payment of $500 received", time: "2 hours ago", read: false },
  ]);

  const [isProfileSheetOpen, setIsProfileSheetOpen] = useState(false);

  const unreadNotifications = notifications.filter(n => !n.read).length;

  // Profile menu options
  const profileMenuItems = [
    { icon: <User className="h-4 w-4" />, label: "Profile", href: "/profile" },
    { icon: <Settings className="h-4 w-4" />, label: "Settings", href: "/settings" },
    { icon: <CreditCard className="h-4 w-4" />, label: "Billing", href: "/billing" },
    { icon: <Users className="h-4 w-4" />, label: "Team", href: "/team" },
    { icon: <BarChart className="h-4 w-4" />, label: "Analytics", href: "/analytics" },
    { icon: <Shield className="h-4 w-4" />, label: "Security", href: "/security" },
    { icon: <Mail className="h-4 w-4" />, label: "Inbox", href: "/inbox" },
    { icon: <Calendar className="h-4 w-4" />, label: "Calendar", href: "/calendar" },
    { icon: <Download className="h-4 w-4" />, label: "Downloads", href: "/downloads" },
    { icon: <LogOut className="h-4 w-4" />, label: "Log Out", href: "/logout", destructive: true },
  ];

  return (
    <header className="sticky top-0 z-50 flex h-12 items-center justify-between border-b bg-background px-4 md:px-6">
      {/* Left Section - Logo */}
      <div className="flex items-center gap-4">
        <Link to="/" className="flex items-center gap-2">
          <div className="h-8 w-8 rounded-lg bg-primary flex items-center justify-center">
            <span className="font-bold text-primary-foreground">L</span>
          </div>
          <span className="hidden text-xl font-bold md:inline-block">Logo</span>
        </Link>

        {/* Username Display - Desktop Only */}
       <div className="hidden lg:flex items-center gap-2">
  <Separator orientation="vertical" className="h-6" />
  <div className="w-15 h-5.5 rounded-2xl p-0 bg-muted flex items-center justify-center">
    <span className="text-[12px] leading-none">username</span>
  </div>
</div>

      </div>

      {/* Right Section - Actions */}
      <div className="flex items-center gap-2">
        {/* Theme Toggle */}
        <ClickDropdownModeToggle />

        <Separator orientation="vertical" className="h-6 hidden md:flex" />

        {/* Chat Icon - Always Visible */}
        <Button variant="ghost" size="icon" className="hover:bg-transparent" asChild>
          <Link to="/chat">
            <MessageSquare className="h-5 w-5" />
            <span className="sr-only">Chat</span>
          </Link>
        </Button>

        {/* Notifications - Desktop Dropdown */}
        <div className="hidden md:block">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" size="icon" className="relative">
                <Bell className="h-5 w-5" />
                {unreadNotifications > 0 && (
                  <Badge
                    variant="destructive"
                    className="absolute -right-1 -top-1 h-5 w-5 rounded-full p-0 text-xs flex items-center justify-center"
                  >
                    {unreadNotifications}
                  </Badge>
                )}
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-80">
              <DropdownMenuLabel className="flex items-center justify-between">
                <span>Notifications</span>
                <Button variant="ghost" size="sm" className="h-auto p-0 text-xs">
                  Mark all as read
                </Button>
              </DropdownMenuLabel>
              <DropdownMenuSeparator />
              <div className="max-h-96 overflow-auto">
                {notifications.map((notification) => (
                  <div key={notification.id}>
                    <DropdownMenuItem className="flex flex-col items-start cursor-pointer p-3">
                      <div className="flex w-full items-start justify-between">
                        <span className="font-medium">{notification.title}</span>
                        {!notification.read && (
                          <div className="h-2 w-2 rounded-full bg-primary" />
                        )}
                      </div>
                      <p className="text-sm text-muted-foreground">
                        {notification.description}
                      </p>
                      <span className="mt-1 text-xs text-muted-foreground">
                        {notification.time}
                      </span>
                    </DropdownMenuItem>
                    <DropdownMenuSeparator />
                  </div>
                ))}
              </div>
              <DropdownMenuItem className="cursor-pointer justify-center text-center">
                View all notifications
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        {/* Notifications - Mobile Link */}
        <Button variant="ghost" size="icon" className="md:hidden relative" asChild>
          <Link to="/notifications">
            <Bell className="h-5 w-5" />
            {unreadNotifications > 0 && (
              <Badge
                variant="destructive"
                className="absolute -right-1 -top-1 h-5 w-5 rounded-full p-0 text-xs flex items-center justify-center"
              >
                {unreadNotifications}
              </Badge>
            )}
          </Link>
        </Button>

        {/* Profile - Desktop Dropdown (no chevron, no hover background) */}
        <div className="hidden md:block">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="ghost" className="p-0 hover:bg-transparent">
                <Avatar className="h-8 w-8">
                  <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                  <AvatarFallback>JD</AvatarFallback>
                </Avatar>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-56">
              <DropdownMenuLabel className="flex flex-col gap-1">
                <span className="font-semibold">John Doe</span>
                <span className="text-xs text-muted-foreground">Administrator</span>
              </DropdownMenuLabel>
              <DropdownMenuSeparator />
              {profileMenuItems.slice(0, 8).map((item) => (
                <DropdownMenuItem key={item.label} className="cursor-pointer" asChild>
                  <Link to={item.href} className="flex items-center">
                    {item.icon}
                    <span className="ml-2">{item.label}</span>
                  </Link>
                </DropdownMenuItem>
              ))}
              <DropdownMenuSeparator />
              <DropdownMenuItem className="cursor-pointer" asChild>
                <Link to={profileMenuItems[9].href} className="flex items-center text-destructive">
                  {profileMenuItems[9].icon}
                  <span className="ml-2">{profileMenuItems[9].label}</span>
                </Link>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>

        {/* Profile - Mobile Sheet (shows avatar image instead of menu icon) */}
        <Sheet open={isProfileSheetOpen} onOpenChange={setIsProfileSheetOpen}>
          <SheetTrigger asChild className="md:hidden">
            <Button variant="ghost" size="icon" className="p-0 hover:bg-transparent">
              <Avatar className="h-8 w-8">
                <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                <AvatarFallback>JD</AvatarFallback>
              </Avatar>
            </Button>
          </SheetTrigger>
          <SheetContent side="right" className="w-full sm:max-w-md">
            <SheetHeader className="border-b pb-4">
              <div className="flex items-center gap-3">
                <Avatar className="h-10 w-10">
                  <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
                  <AvatarFallback>JD</AvatarFallback>
                </Avatar>
                <div>
                  <SheetTitle>John Doe</SheetTitle>
                  <p className="text-sm text-muted-foreground">Administrator</p>
                </div>
              </div>
            </SheetHeader>
            
            <div className="py-4">
              <div className="space-y-1">
                {profileMenuItems.map((item) => (
                  <Button
                    key={item.label}
                    variant="ghost"
                    className={`w-full justify-start ${item.destructive ? 'text-destructive hover:text-destructive' : ''}`}
                    asChild
                    onClick={() => setIsProfileSheetOpen(false)}
                  >
                    <Link to={item.href} className="flex items-center">
                      {item.icon}
                      <span className="ml-3">{item.label}</span>
                    </Link>
                  </Button>
                ))}
              </div>
            </div>
            
            <div className="absolute bottom-0 left-0 right-0 border-t p-4">
              <div className="flex items-center justify-between">
                <span className="text-sm text-muted-foreground">Theme</span>
                <HorizontalModeToggle/>
              </div>
            </div>
          </SheetContent>
        </Sheet>
      </div>
    </header>
  );
}