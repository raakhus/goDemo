import React from 'react'
import { Navigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'

export const PrivateRoute: React.FC<{children: JSX.Element}> = ({ children }) => {
  const { user } = useAuth()
  if (user === null) {
    return <Navigate to="/login" replace />
  }
  return children
}
