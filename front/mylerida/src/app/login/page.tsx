"use client"; // Se requiere para componentes interactivos en Next.js 13

import React, { useState } from "react";
import { HttpAuthAdapter } from "@/infrastructure/adapter/authAdapter";
import { LoginUserUseCase } from "@/application/usecaselogin/loginuser";
import styles from "../login/login.module.css";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [token, setToken] = useState("");

  // Lee la variable de entorno para la URL base
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:60000";
  const authAdapter = new HttpAuthAdapter(baseURL);
  const loginUseCase = new LoginUserUseCase(authAdapter);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    setError("");

    try {
      const accessToken = await loginUseCase.execute(email, password);
      setToken(accessToken);
      // Aquí podrías guardar el token en cookies o localStorage para su uso posterior
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else if (err instanceof TypeError) {
        setError("Network error: Failed to fetch");
      } else if (typeof err === "string") {
        setError(err);
      } else {
        setError("An unexpected error occurred");
      }
    }
  }

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Login</h1>
      <form onSubmit={handleSubmit} className={styles.form}>
        <div>
          <label className={styles.label}>Email:</label><br/>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className={`${styles.input}`}
            required
          />
        </div>
        <div style={{ marginTop: 10 }}>
          <label className={styles.label}>Password:</label><br/>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className={`${styles.input}`}
            required
          />
        </div>
        <button type="submit" className={styles.button}>Iniciar sesión</button>
      </form>
      {error && <p style={{ color: "red" }}>Error: {error}</p>}
      {token && <p style={{ color: "green" }}>Token: {token}</p>}
    </div>
  );
}
