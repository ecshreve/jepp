import React from 'react';
import './Button.css';
import { ButtonProps } from 'react-bootstrap';

interface Props extends ButtonProps {
  /**
   * Is this the principal call to action on the page?
   */
  primary?: boolean;
  /**
   * What background color to use
   */
  backgroundColor?: string;
  /**
   * How large should the button be?
   */
  size?: 'sm' | 'lg';
  /**
   * Button contents
   */
  label: string;
  /**
   * Optional click handler
   */
  onClick?: () => void;
}

/**
 * Primary UI component for user interaction
 */
export const Button = ({
  primary = false,
  size = 'sm',
  backgroundColor,
  label,
  ...props
}: Props) => {
  const mode = primary ? 'jepp-button--primary' : 'jepp-button--secondary';
  return (
    <button
      type="button"
      className={['jepp-button', `jepp-button--${size}`, mode].join(' ')}
      style={{ backgroundColor }}
      {...props}
    >
      {label}
    </button>
  );
};