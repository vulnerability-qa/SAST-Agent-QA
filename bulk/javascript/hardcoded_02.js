// CWE-798: Hardcoded payment API key (test value — for SAST detection demo)
const PAYMENT_API_KEY = 'test_hardcoded_key_12345_never_commit_real_keys';
const paymentClient = { key: PAYMENT_API_KEY };
module.exports = paymentClient;
