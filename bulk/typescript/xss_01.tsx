// CWE-79: XSS via dangerouslySetInnerHTML in React
import React from 'react';
interface Props { html: string; }
export const RichText: React.FC<Props> = ({ html }) => (
  <div dangerouslySetInnerHTML={{ __html: html }} />
);
