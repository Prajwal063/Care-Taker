import express from 'express';
import passport from 'passport';
import { CLIENT_URL } from '../env';

const router = express.Router();

router.route('/login/success').get((req, res) => {
    if (req.user) {
        res.status(200).json({
            error: false,
            message: 'Successfully Logged In',
            user: req.user,
        });
    } else {
        res.status(403).json({ error: true, message: 'Not Authorized' });
    }
});

router.route('/login/failed').get((req, res) => {
    res.status(401).json({
        error: true,
        message: 'Failed to authenticate with Google',
    });
});

router
    .route('/google')
    .get(passport.authenticate('google', { scope: ['profile', 'email'] }));

router.route('/google/callback').get(
    passport.authenticate('google', {
        successRedirect: CLIENT_URL,
        failureRedirect: '/auth/login/failed',
    })
);

router.route('/logout').get((req, res) => {
    req.logout({ keepSessionInfo: false }, () => {});
    res.redirect(CLIENT_URL);
});

export default router;
