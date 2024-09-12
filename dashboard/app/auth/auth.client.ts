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


// WorkOS and JWT stuff

// export const workos = new WorkOS(process.env.WORKOS_API_KEY);

// export function getJwtSecretKey() {
//   const secret = process.env.JWT_SECRET_KEY;

//   if (!secret) {
//     throw new Error('JWT_SECRET_KEY is not set');
//   }

//   return new Uint8Array(Buffer.from(secret, 'base64'));
// }

// export async function verifyJwtToken(token: string) {
//   try {
//     const { payload } = await jwtVerify(token, getJwtSecretKey());
//     return payload;
//   } catch (error) {
//     return null;
//   }
// }

// export async function getAccessToken(): Promise<string | undefined> {
//   return cookies().get('token')?.value;
// }