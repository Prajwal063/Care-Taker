import mongoose from 'mongoose';

const ServiceSchema = new mongoose.Schema(
    {
        title: {
            type: String,
            required: true,
            unique: true,
            min: 3,
            max: 32,
            trim: true,
        },
        description: {
            type: String,
            required: true,
            trim: true,
            min: 3,
            max: 500,
        },
        email: {
            type: String,
            required: true,
            trim: true,
        },
        phone: {
            type: String,
            required: true,
            trim: true,
            min: 10,
            max: 10,
        },
        address: {
            type: String,
            required: true,
            trim: true,
            max: 100,
        },
        picture: {
            type: String,
            required: true,
            trim: true,
        },
    },
    { timestamps: true }
);

ServiceSchema.virtual('id').get(function () {
    return this._id.toHexString();
});

// To ensure virtual fields are serialized.
ServiceSchema.set('toJSON', {
    virtuals: true,
    transform(doc, ret, options) {
        delete ret._id;
        delete ret.__v;
        return ret;
    },
});

export default mongoose.model('Service', ServiceSchema);
