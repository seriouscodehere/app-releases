import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";
import type { LinkProps } from "react-router-dom";
import { cn } from "@/lib/utils";

interface LinkButtonProps extends React.ComponentProps<typeof Button> {
  to: LinkProps["to"];
  linkProps?: Omit<LinkProps, "to" | "children">;
}

export function LinkButton({
  to,
  linkProps,
  className,
  children,
  ...buttonProps
}: LinkButtonProps) {
  return (
    <Button
      className={cn("", className)}
      asChild
      {...buttonProps}
    >
      <Link to={to} {...linkProps}>
        {children}
      </Link>
    </Button>
  );
}