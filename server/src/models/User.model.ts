import mongoose from 'mongoose';

const UserSchema = new mongoose.Schema(
    {
        name: {
            type: String,
            required: true,
            min: 3,
            max: 32,
            trim: true,
        },
        email: {
            type: String,
            required: true,
            unique: true,
            trim: true,
        },
        googleId: {
            type: String,
            required: true,
            unique: true,
            trim: true,
        },
        picture: {
            type: String,
            required: true,
            trim: true,
        },
    },
    { timestamps: true }
);

UserSchema.virtual('id').get(function () {
    return this._id.toHexString();
});

// To ensure virtual fields are serialized.
UserSchema.set('toJSON', {
    virtuals: true,
    transform(doc, ret, options) {
        delete ret._id;
        delete ret.__v;
        delete ret.createdAt;
        delete ret.updatedAt;
        return ret;
    },
});

export default mongoose.model('User', UserSchema);
