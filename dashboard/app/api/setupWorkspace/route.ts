import { validateRequest } from '@/app/auth/validate';
import { getNewRootKey } from '@/lib/apiKeyService/rootKey';
import { GetOrCreateDefaultWorkspaceForUser } from '@/lib/apiKeyService/workspace';
import { NextRequest, NextResponse } from 'next/server';
import { sealData } from 'iron-session';
import { cookies } from 'next/headers';

export async function GET(_: NextRequest) {
  const { user, session } = await validateRequest();

  const workspaceId = await GetOrCreateDefaultWorkspaceForUser(user!.githubId, session!.id);

  const rootKey = await getNewRootKey(workspaceId, `${user!.id}-default-root-key`);
    const encryptedRootKey = await sealData(rootKey, {
      password: process.env.ROOTKEY_PASSWORD as string,
    });

    cookies().set("apikeyper-auth", encryptedRootKey);

  return NextResponse.json({ message: "success", workspaceId });
}