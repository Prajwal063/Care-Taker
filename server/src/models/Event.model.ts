import mongoose from 'mongoose';

const EventStatus = ['UPCOMING', 'ONGOING', 'COMPLETED', 'CANCELLED'];

const EventSchema = new mongoose.Schema(
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
            unique: true,
            trim: true,
            min: 3,
            max: 500,
        },
        email: {
            type: String,
            required: true,
            unique: true,
            trim: true,
        },
        phone: {
            type: String,
            required: true,
            unique: true,
            trim: true,
            min: 10,
            max: 10,
        },
        address: {
            type: String,
            required: true,
            unique: true,
            trim: true,
            max: 100,
        },
        picture: {
            type: String,
            required: true,
            trim: true,
        },
        date: {
            type: Date,
            required: true,
        },
        time: {
            type: String,
            required: true,
        },
        price: {
            type: Number,
            required: true,
        },
        status: {
            type: String,
            enum: EventStatus,
            required: true,
        },
        registeredUsers: {
            type: [String],
            required: true,
            default: [],
        },
    },
    { timestamps: true }
);

EventSchema.virtual('id').get(function () {
    return this._id.toHexString();
});

// To ensure virtual fields are serialized.
EventSchema.set('toJSON', {
    virtuals: true,
    transform(doc, ret, options) {
        delete ret._id;
        delete ret.__v;
        return ret;
    },
});

export default mongoose.model('Event', EventSchema);
