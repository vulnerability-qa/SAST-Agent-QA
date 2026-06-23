// CWE-915: Mass Assignment
const User = require('./models/User');
async function updateProfile(req, res) {
  await User.findByIdAndUpdate(req.user.id, req.body); // isAdmin settable via body
  res.json({ ok: true });
}
