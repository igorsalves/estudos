import { Injectable } from '@nestjs/common';
import { InjectConnection, InjectModel } from '@nestjs/mongoose';
import mongoose, { Model } from 'mongoose';
import { Asset } from '../assets/entities/asset.entity';
import { CreateWalletDto } from './dto/create-wallet.dto';
import { WalletAsset } from './entities/wallet-asset.entity';
import { Wallet } from './entities/wallet.entity';

@Injectable()
export class WalletsService {
  constructor(
    @InjectModel(Wallet.name) private readonly walletSchema: Model<Wallet>,
    @InjectModel(WalletAsset.name)
    private readonly walletAssetSchema: Model<WalletAsset>,
    @InjectConnection() private readonly connection: mongoose.Connection,
  ) {}

  create(createWalletDto: CreateWalletDto) {
    return this.walletSchema.create(createWalletDto);
  }

  findAll() {
    return this.walletSchema.find();
  }

  findOne(id: string) {
    return this.walletSchema.findById(id).populate([
      {
        path: 'assets', //walletasset
        populate: ['asset'],
      },
    ]) as Promise<
      (Wallet & { assets: (WalletAsset & { asset: Asset })[] }) | null
    >;
  }

  async createWalletAsset(data: {
    walletId: string;
    assetId: string;
    shares: number;
  }) {
    const session = await this.connection.startSession();
    // eslint-disable-next-line @typescript-eslint/await-thenable
    await session.startTransaction();

    try {
      const docs = await this.walletAssetSchema.create(
        [
          {
            wallet: data.walletId,
            asset: data.assetId,
            shares: data.shares,
          },
        ],
        { session },
      );
      const walletAsset = docs[0];
      await this.walletSchema.updateOne(
        { _id: data.walletId },
        { $push: { assets: walletAsset._id } },
        { session },
      );

      await session.commitTransaction();
      return walletAsset;
    } catch (error) {
      console.error(error);
      await session.abortTransaction();
      throw error;
    } finally {
      await session.endSession();
    }
  }
}
