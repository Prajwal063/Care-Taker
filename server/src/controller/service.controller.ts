import type { NextFunction, Request, Response } from 'express';
import AppError from '../helper/AppError';
import { Service } from '../models';

export const getAllServices = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { search } = req.query;

        const services = await Service.find({
            title: { $regex: search || '', $options: 'i' },
        });

        res.status(200).json({
            message: 'Fetched services Successfully!',
            data: services,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};

export const createService = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const service = new Service(req.body);

        await service.save();

        res.status(201).json({
            message: 'Service created Successfully!',
            data: service,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};

export const getServiceById = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        const service = await Service.findById(id);

        res.status(201).json({
            message: 'Fetched service Successfully!',
            data: service,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};

export const updateService = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        const service = await Service.findByIdAndUpdate(id, req.body, {
            new: true,
        });

        res.status(200).json({
            message: 'Service updated Successfully!',
            data: service,
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};

export const deleteService = async (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    try {
        const { id } = req.params;

        const service = await Service.findByIdAndDelete(id);
        if (!service) {
            throw new AppError({
                message: 'Service not found!',
                statusCode: 404,
            });
        }

        res.status(200).json({
            message: 'Service deleted Successfully!',
        });
    } catch (err: any) {
        next(
            new AppError({
                message: err.message || 'Server error occurred!',
                statusCode: err.statusCode || 400,
                stack: err.stack || '',
            })
        );
    }
};
