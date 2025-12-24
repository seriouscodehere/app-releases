import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Table, TableBody, TableCell, TableRow } from "@/components/ui/table"
import { Badge } from "@/components/ui/badge"
import { GlobalTopbar } from "@/components/reusable/navbars/global_topbar"

export default function UserProfile() {
  return (
    <>
    <GlobalTopbar/>
    <div className="w-full max-w-5xl mx-auto space-y-6">
      {/* Cover and Profile */}
      <Card className="overflow-hidden">
        <div className="h-48 w-full relative">
          <img
            src="https://media.istockphoto.com/id/183765619/photo/green-fields-and-mounatins.jpg?s=612x612&w=0&k=20&c=bxHFYIZQdrwkgeKExOPrDG8j_RgLTeCwB3OywWOBv-A="
            alt="Cover"
            className="h-full w-full object-cover"
          />
          <div className="absolute left-6 -bottom-10">
            <Avatar className="h-24 w-24 border-4 border-background">
              <AvatarImage src="https://plus.unsplash.com/premium_photo-1689568126014-06fea9d5d341" />
              <AvatarFallback>JD</AvatarFallback>
            </Avatar>
          </div>
        </div>

        <CardContent className="pt-14">
          <h2 className="text-2xl font-semibold">John Doe</h2>
          <p className="text-sm text-muted-foreground">@johndoe</p>
          <div className="mt-2 flex gap-2">
            <Badge>Active</Badge>
            <Badge variant="secondary">User</Badge>
          </div>
        </CardContent>
      </Card>

      {/* Tabs */}
      <Tabs defaultValue="personal" className="w-full">
        <TabsList className="flex overflow-x-auto whitespace-nowrap p-0 border-b border-border scrollbar-hide">
          {[
            ["personal", "Personal Details"],
            ["posts", "Posts"],
            ["jobs", "Jobs"],
            ["applied", "Applied Jobs"],
          ].map(([value, label]) => (
            <TabsTrigger
              key={value}
              value={value}
              className="px-4 py-3 text-sm rounded-none border-b-2 border-transparent data-[state=active]:border-primary data-[state=active]:text-primary shrink-0 hover:bg-muted/50 transition-colors"
            >
              {label}
            </TabsTrigger>
          ))}
        </TabsList>

        {/* Personal Details Tab */}
        <TabsContent value="personal" className="mt-6">
          <Card>
            <CardHeader>
              <CardTitle>Personal Details</CardTitle>
            </CardHeader>
            <CardContent>
              <Table>
                <TableBody>
                  <TableRow>
                    <TableCell className="font-medium">Full Name</TableCell>
                    <TableCell>John Michael Doe</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell className="font-medium">Username</TableCell>
                    <TableCell>johndoe</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell className="font-medium">Email</TableCell>
                    <TableCell>john@example.com</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell className="font-medium">Status</TableCell>
                    <TableCell>
                      <Badge>Active</Badge>
                    </TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell className="font-medium">Created At</TableCell>
                    <TableCell>2023-04-12</TableCell>
                  </TableRow>
                  <TableRow>
                    <TableCell className="font-medium">Last Updated</TableCell>
                    <TableCell>2025-01-08</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </CardContent>
          </Card>
        </TabsContent>

        {/* Other Tabs */}
        {[
          ["posts", "Posts content here"],
          ["jobs", "Jobs content here"],
          ["applied", "Applied Jobs content here"],
          ["recommendations", "Recommendations content here"],
        ].map(([value, text]) => (
          <TabsContent key={value} value={value} className="mt-6">
            <Card>
              <CardHeader>
                <CardTitle>{value.charAt(0).toUpperCase() + value.slice(1)}</CardTitle>
              </CardHeader>
              <CardContent>{text}</CardContent>
            </Card>
          </TabsContent>
        ))}
      </Tabs>
    </div>
    </>

  )
}
