import { cookies } from 'next/headers';
import { unsealData } from 'iron-session';

export async function getRootKey(): Promise<string | undefined> {
  const key = cookies().get('apikeyper-auth')?.value;
  if (key) {
    return unsealData(key, {
      password: process.env.ROOTKEY_PASSWORD as string,
    });
  }
}