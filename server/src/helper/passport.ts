import passport from 'passport';
import GoogleStrategy from 'passport-google-oauth20';
import { GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET } from '../env';
import { User } from '../models';
import { handlePromise } from '../utils';

export default (function () {
    passport.use(
        new GoogleStrategy.Strategy(
            {
                clientID: GOOGLE_CLIENT_ID,
                clientSecret: GOOGLE_CLIENT_SECRET,
                callbackURL: '/auth/google/callback',
                scope: ['profile', 'email'],
            },
            async function (accessToken, refreshToken, profile, callback) {
                console.log(profile);

                // User.findOneAndUpdate(
                //     { googleId: profile.id },
                //     {
                //         $setOnInsert: {
                //             googleId: profile.id,
                //             name: profile.displayName,
                //             email: profile?.emails[0].value ?? '',
                //             picture: profile?.photos[0].value ?? '',
                //         },
                //     },
                //     { upsert: true, returnNewDocument: true }
                // );

                let [user, _] = await handlePromise(
                    User.findOne({ googleId: profile.id })
                );
                if (!user) {
                    let user = new User({
                        googleId: profile.id,
                        name: profile.displayName,
                        email: profile?.emails[0].value ?? '',
                        picture: profile?.photos[0].value ?? '',
                    });
                    await user.save();
                }

                callback(null, profile);
            }
        )
    );

    passport.serializeUser((user, done) => {
        done(null, user);
    });

    passport.deserializeUser(async (user, done) => {
        console.log('Deserialize User: ', user);
        const [usr, error] = await handlePromise(
            User.findOne({ googleId: user.id })
        );

        done(error, usr ?? null);
    });
});
