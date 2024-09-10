import Image from "next/image";

const mockUser =  {
  id: "1234567890",
  firstName: "John",
  lastName: "Doe",
  email: "john.doe@example.com",
  profilePictureUrl: "https://lh3.googleusercontent.com/a/ACg8ocJeui0JxrWCISlo4nWnCmANIhvlOzFgq-dQcRGe03W8TK7ZxN5MeQ=s317-c-no",
}

export default function UserItem() {
  return (
    <div className="flex items-center justify-between gap-2 rounded-[8px] p-1 text-black/50 dark:text-white">
      <Image src={mockUser.profilePictureUrl!} alt="" width="50" height="50" className="rounded-full" />
      <div className="grow">
        <p className="text-[16px] font-bold">{mockUser?.firstName}</p>
        <p className="text-[12px] text-neutral-500">
          {mockUser?.email}
        </p>
      </div>
    </div>
  );
}
