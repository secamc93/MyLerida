// src/application/use-cases/LoginUserUseCase.ts

import { AuthPort } from "@/domain/ports/auth";

export class LoginUserUseCase {
  constructor(private readonly authPort: AuthPort) {}

  public async execute(email: string, password: string): Promise<string> {
    // Aquí podrías agregar validaciones o lógica adicional si fuera necesario.
    return await this.authPort.login(email, password);
  }
}
