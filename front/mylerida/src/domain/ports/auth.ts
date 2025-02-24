// src/application/ports/AuthPort.ts

export interface AuthPort {
    login(email: string, password: string): Promise<string>;
  }
  