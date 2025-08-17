import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'

export default function Login() {
  const [email, setEmail] = useState('test@example.com')
  const [password, setPassword] = useState('Passw0rd!')
  const [err, setErr] = useState<string | null>(null)
  const nav = useNavigate()
  const { login } = useAuth()

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault()
    setErr(null)
    try {
      await login(email, password)
      nav('/protected', { replace: true })
    } catch (e: any) {
      setErr(e.message || 'Login failed')
    }
  }

  return (
    <div style={{ maxWidth: 360, margin: '6rem auto', fontFamily: 'system-ui' }}>
      <h1>Sign in</h1>
      <form onSubmit={onSubmit}>
        <label>Email<br/>
          <input value={email} onChange={e => setEmail(e.target.value)} />
        </label><br/><br/>
        <label>Password<br/>
          <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
        </label><br/><br/>
        <button type="submit">Sign in</button>
      </form>
      {err && <p style={{ color: 'crimson' }}>{err}</p>}
      <p style={{ marginTop: 12, color: '#555' }}>Try test@example.com / Passw0rd!</p>
    </div>
  )
}
