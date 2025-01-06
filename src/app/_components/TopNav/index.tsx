import { SignedIn, SignedOut, SignInButton, UserButton } from "@clerk/nextjs";
import Link from "next/link";
import Image from "next/image";

const bannerPath = "/banner.png";

export default function TopNav() {
  return (
    <nav className="sticky top-0 flex w-full items-center justify-between border-b border-primary bg-black p-4 text-xl font-semibold">
      <Link href="/" className="h-auto w-auto">
        <Image
          src={bannerPath}
          alt="Logo"
          width={400}
          height={0}
          priority={true}
          className="h-full w-auto object-contain"
        />
      </Link>
      <div>
        <SignedOut>
          <SignInButton />
        </SignedOut>
        <SignedIn>
          <UserButton />
        </SignedIn>
      </div>
    </nav>
  );
}
