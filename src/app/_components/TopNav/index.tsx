import { SignedIn, SignedOut, SignInButton, UserButton } from "@clerk/nextjs";
import Link from "next/link";
import Image from "next/image";

export default function TopNav() {
  return (
    <nav className="sticky top-0 flex w-full items-center justify-between border-b border-primary bg-black p-4 text-xl font-semibold">
      <Link href="/">
        <Image
          src="/banner.png"
          alt="Logo"
          width={400}
          height={0}
          layout="intrinsic"
        />
      </Link>
      <div className="">
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
