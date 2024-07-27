import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
export declare class Figure extends UsxItemContainer<Text> {
    style: string;
    alt?: string;
    file: string;
    size?: string;
    loc?: string;
    copy?: string;
    ref?: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class FigureFactory extends UsxItemFactory<Figure> {
    static readonly instance: FigureFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Figure;
}
//# sourceMappingURL=figure.d.ts.map