// src/infrastructure/adapters/HttpAuthAdapter.ts

import { AuthPort } from "@/domain/ports/auth";

export class HttpAuthAdapter implements AuthPort {
  private readonly baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  public async login(email: string, password: string): Promise<string> {
    const response = await fetch(`${this.baseURL}/api/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || `Error en login: ${response.statusText}`);
    }

    const data = await response.json();
    // Suponemos que data tiene la estructura { access_token: string }
    return data.access_token;
  }
}
