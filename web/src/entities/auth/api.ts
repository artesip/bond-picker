import type { LoginData, RegistrationData, User } from './model';

const headers = new Headers({
  'Content-Type': 'application/json',
});

export async function Login(data: LoginData) {
  const resp = await fetch('/backend/api/v1/auth/login', 
    {
      method : 'POST',
      body   : JSON.stringify(data),
      headers: headers,
    }
  );

  if (resp.status === 403) {
    throw Error('Неверный логин или пароль');
  }

  if (!resp.ok) {
    throw Error('Ошибка входа');
  }
}

export async function Registration(data: RegistrationData) {
  const resp = await fetch('/backend/api/v1/auth/registration', 
    {
      method : 'POST',
      body   : JSON.stringify(data),
      headers: headers,
    }
  );

  if (resp.status === 409) {
    throw Error('Пользователь с текущим login уже есть');
  }

  if (!resp.ok) {
    throw Error('Ошибка регистрации');
  }
}

export async function Logout() {
  const resp = await fetch('/backend/api/v1/auth/logout', 
    {
      method : 'POST',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка выхода');
  }
}

export async function Me(): Promise<User> {
  const resp = await fetch('/backend/api/v1/auth/me', 
    {
      method : 'GET',
      headers: headers,
    }
  );

  if (!resp.ok) {
    throw Error('Ошибка выхода');
  }

  return await resp.json();
}