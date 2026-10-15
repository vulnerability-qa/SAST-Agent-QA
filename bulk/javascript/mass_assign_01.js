// CWE-915: Mass Assignment
const User = require('./models/User');
async function updateProfile(req, res) {
  const allowedFields = ['name', 'email', 'bio', 'phone', 'preferences'];
  const allowedUpdates = {};
  allowedFields.forEach(field => {
    if (req.body[field] !== undefined) {
      allowedUpdates[field] = req.body[field];
    }
  });
  await User.findByIdAndUpdate(req.user.id, allowedUpdates);
  res.json({ ok: true });
}
