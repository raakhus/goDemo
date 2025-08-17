import React, { useEffect, useState } from 'react'
import { useAuth } from '../context/AuthContext'
import { api } from '../lib/api'

export default function Protected() {
  const { user, logout } = useAuth()
  const [msg, setMsg] = useState<string>('')

  useEffect(() => {
    api<{message:string}>('/api/secret').then(r => setMsg(r.message)).catch(() => setMsg(''))
  }, [])

  return (
    <div style={{ maxWidth: 480, margin: '6rem auto', fontFamily: 'system-ui' }}>
      <h1>{msg || 'You made it.'}</h1>
      <p>Signed in as <b>{user?.email}</b></p>
      <button onClick={logout}>Sign out</button>
    </div>
  )
}
