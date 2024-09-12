import { NextRequest, NextResponse } from 'next/server';
import { createApi } from '@/lib/apiKeyService/api';

export async function POST(request: NextRequest) {
  const req = await request.json();
  await createApi({
    apiName: req.apiName,
  });
  return NextResponse.json({ message: "success" });
}