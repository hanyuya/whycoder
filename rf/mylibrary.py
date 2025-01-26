def returnlist():
    return [i for i in range(10)]


def return_dict():
    return {"a": "hahhahahaahah"}


# 以下划线开头的函数不能作为RF关键字
def _returnlist2():
    return [1, 2]