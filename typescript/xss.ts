import express, { Request, Response } from 'express';

const app = express();

const escapeHtml = (str: string): string => {
  return str.replace(/[&<>"']/g, (char) => {
    const entities: { [key: string]: string } = {
      '&': '&amp;',
      '<': '&lt;',
      '>': '&gt;',
      '"': '&quot;',
      "'": '&#39;'
    };
    return entities[char];
  });
};

app.get('/greet', (req: Request, res: Response) => {
  const name = req.query.name as string || '';
  // CWE-79: user input reflected directly into HTML without escaping.
  res.send(`<h1>Hello, ${escapeHtml(name)}!</h1>`);
});

app.listen(3000);
