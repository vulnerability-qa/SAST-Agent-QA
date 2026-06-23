// CWE-79: XSS via dangerouslySetInnerHTML in React
import React from 'react';
import DOMPurify from 'dompurify';
interface Props { html: string; }
export const RichText: React.FC<Props> = ({ html }) => {
  const sanitizedHtml = DOMPurify.sanitize(html);
  return <div dangerouslySetInnerHTML={{ __html: sanitizedHtml }} />;
};
